package services

import (
	"bifrost/extensions"
	"bifrost/helpers"
	"bifrost/models/media"
	"bifrost/models/post"
	"bifrost/models/post/payloads"
	"bifrost/models/post/utils"
	global_shared "bifrost/models/shared"

	form "github.com/go-playground/form/v4"

	"bifrost/models/user"
	"bifrost/repositories"
	"bifrost/types"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type PostService struct {
	mediaRepo *repositories.MediaRepository
	userRepo  *repositories.UserRepository
	postRepo  *repositories.PostRepository
}

func NewPostService(
	userRepo *repositories.UserRepository,
	postRepo *repositories.PostRepository,
	mediaRepo *repositories.MediaRepository) *PostService {
	return &PostService{postRepo: postRepo, mediaRepo: mediaRepo, userRepo: userRepo}
}

func (s *PostService) CreatePost(request map[string][]string, files []*multipart.FileHeader, author *user.User) (*post.Post, error) {
	fmt.Println("POST_SERVICE:CreatePost")

	type PollForm struct {
		ID       string   `form:"id"`
		Question string   `form:"question"`
		Duration string   `form:"duration"`
		Options  []string `form:"options"` // options[] → slice
	}

	//Dot Notation

	type PostForm struct {
		// Temel post bilgileri
		ParentId string   `form:"parentPostId"`
		Title    string   `form:"title"`
		Summary  string   `form:"summary"`
		Content  string   `form:"content"`
		Audience string   `form:"audience"`
		Hashtags []string `form:"hashtags[]"` // body[hashtags][0], body[hashtags][1]...

		Polls []PollForm `form:"polls"`

		// Event bilgileri
		EventTitle       string `form:"event[title]"`
		EventDescription string `form:"event[description]"`
		EventDate        string `form:"event[date]"` // YYYY-MM-DD
		EventTime        string `form:"event[time]"` // HH:MM

		// Location bilgileri
		LocationAddress string  `form:"location[address]"`
		LocationLat     float64 `form:"location[lat]"`
		LocationLng     float64 `form:"location[lng]"`
	}
	decoder := form.NewDecoder()
	postForm := PostForm{}

	if err := decoder.Decode(&postForm, request); err != nil {
		fmt.Println("Form decode error:", err)
		return nil, err
	}

	tx := s.postRepo.DB().Begin() // transaction başlat
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	node, err := helpers.NewNode(1)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create snowflake node: %w", err)
	}
	var parentUUID *uuid.UUID

	if len(postForm.ParentId) > 0 {
		parsed, err := uuid.Parse(postForm.ParentId)
		if err != nil {
			// Hata yönetimi: geçersiz UUID
			fmt.Println("Invalid ParentId:", err)
			// İster panic, ister return, ister loglayabilirsin
			parentUUID = nil
		} else {
			parentUUID = &parsed
		}
	}

	defaultLanguage := "en"
	newPost := &post.Post{
		ID:        uuid.New(),
		ParentID:  parentUUID,
		AuthorID:  author.ID,
		Published: false,
		Type:      post.PostTypeTimeline,
		Title:     utils.MakeLocalizedString(defaultLanguage, postForm.Title),
		Content:   utils.MakeLocalizedString(defaultLanguage, postForm.Content),
		Summary:   utils.MakeLocalizedString(defaultLanguage, postForm.Summary),
		PublicID:  node.Generate().Int64(),
	}

	// Post DB'ye ekle
	if err := tx.Create(newPost).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Post media

	for _, f := range files {
		mediaModel, err := s.mediaRepo.AddMedia(tx, newPost.ID, media.OwnerPost, media.RolePost, f)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		newPost.Attachments = append(newPost.Attachments, mediaModel)
	}

	for _, pollInfo := range postForm.Polls {
		poll := &payloads.Poll{
			ID:              uuid.New(),
			ContentableID:   newPost.ID,
			ContentableType: payloads.ContentablePollPost,
			Question:        *utils.MakeLocalizedString(defaultLanguage, pollInfo.Question),
			Duration:        pollInfo.Duration,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		for _, choiceLabel := range pollInfo.Options {
			poll.Choices = append(poll.Choices, payloads.PollChoice{
				ID:        uuid.New(),
				PollID:    poll.ID,
				Label:     *utils.MakeLocalizedString(defaultLanguage, choiceLabel),
				VoteCount: 0,
			})
		}

		if err := s.postRepo.CreatePoll(poll); err != nil {
			tx.Rollback()
			return nil, err
		}

		newPost.Poll = append(newPost.Poll, poll)
	}

	var locationPost *global_shared.Location = nil // varsayılan olarak nil
	var locationPoint *extensions.PostGISPoint = nil

	// location

	if postForm.LocationLat != 0 && postForm.LocationLng != 0 {
		locationPoint = &extensions.PostGISPoint{
			Lat: postForm.LocationLat,
			Lng: postForm.LocationLng,
		}

		locationPost = &global_shared.Location{
			ID:              uuid.New(),
			ContentableType: global_shared.LocationOwnerPost,
			ContentableID:   newPost.ID,
			Address:         &postForm.LocationAddress,
			Latitude:        &postForm.LocationLat,
			Longitude:       &postForm.LocationLng,
			LocationPoint:   locationPoint,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		if err := s.userRepo.UpsertLocation(locationPost); err != nil {
			return nil, err
		}
	}

	//  Event
	var evt *payloads.Event
	if len(postForm.EventTitle) > 0 {
		startTime := time.Time{}
		if len(postForm.EventDate) > 0 {
			if len(postForm.EventTime) > 0 {
				if parsedTime, err := time.Parse("2006-01-02 15:04", postForm.EventDate+" "+postForm.EventTime); err == nil {
					startTime = parsedTime
				}
			}
		}

		evt = &payloads.Event{
			ID:          uuid.New(),
			PostID:      newPost.ID,
			Title:       *utils.MakeLocalizedString(defaultLanguage, postForm.EventTitle),
			Description: *utils.MakeLocalizedString(defaultLanguage, postForm.EventDescription),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			StartTime:   &startTime,
		}

		// Event DB'ye ekle
		if err := tx.Create(evt).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		locationEvent := &global_shared.Location{
			ID:              uuid.New(),
			ContentableType: global_shared.LocationOwnerEvent,
			ContentableID:   evt.ID,
			Address:         &postForm.LocationAddress,
			Latitude:        &postForm.LocationLat,
			Longitude:       &postForm.LocationLng,
			LocationPoint:   locationPoint,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		if err := tx.Create(locationEvent).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		evt.Location = locationEvent
		newPost.Event = evt
	}

	if err := tx.Save(newPost).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	fmt.Println("NEW POST", newPost.ID)

	lastPost, _ := s.postRepo.GetPostByID(newPost.ID)
	return lastPost, nil
}

func (s *PostService) GetPostByID(id uuid.UUID) (*post.Post, error) {
	postData, err := s.postRepo.GetPostByID(id)
	if err != nil {
		return nil, fmt.Errorf("GetPostByID error: %w", err)
	}
	return postData, nil
}

func (s *PostService) GetPostByPublicID(id int64) (*post.Post, error) {
	postData, err := s.postRepo.GetPostByPublicID(id)
	if err != nil {
		return nil, fmt.Errorf("GetPostByID error: %w", err)
	}
	return postData, nil
}

func (s *PostService) GetTimeline(limit int, cursor *int64) (types.TimelineResult, error) {
	// Repo fonksiyonunu çağırıyoruz
	posts, err := s.postRepo.GetTimeline(limit, cursor)
	if err != nil {
		return types.TimelineResult{}, err
	}
	return posts, nil
}
