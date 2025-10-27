package services

import (
	"bifrost/extensions"
	"bifrost/helpers"
	"bifrost/models/media"
	"bifrost/models/post"
	"bifrost/models/post/payloads"
	"bifrost/models/post/utils"
	"bifrost/models/shared"
	global_shared "bifrost/models/shared"
	"bifrost/models/user"
	"bifrost/repositories"
	"fmt"
	"mime/multipart"
	"time"

	form "github.com/go-playground/form/v4"
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

	type PostForm struct {
		// Temel post bilgileri
		Content  string   `form:"content"`
		Audience string   `form:"audience"`
		Hashtags []string `form:"hashtags[]"` // body[hashtags][0], body[hashtags][1]...

		// Poll bilgileri
		PollDuration int      `form:"poll[duration]"`
		PollOptions  []string `form:"poll[options]"` // body[poll][options][0..n]

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

	fmt.Println("REQUEST", postForm.Hashtags)

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

	defaultLanguage := "en"
	newPost := &post.Post{
		ID:        uuid.New(),
		AuthorID:  author.ID,
		Published: false,
		Type:      post.PostTypeTimeline,
		Title:     nil,
		Content:   utils.MakeLocalizedString(defaultLanguage, postForm.Content),
		Summary:   nil,
		PublicID:  node.Generate().Int64(),
	}

	// Post DB'ye ekle
	if err := tx.Create(newPost).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	fmt.Println("postForm", postForm)

	// Post media
	fmt.Println("files", files)

	for _, f := range files {
		mediaModel, err := s.mediaRepo.AddMedia(tx, newPost.ID, media.OwnerPost, media.RolePost, f)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		newPost.Attachments = append(newPost.Attachments, mediaModel)
	}

	// Poll
	hasPoll := len(postForm.PollOptions) > 0

	if hasPoll {
		fmt.Println("SAVE:POLL")
		poll := &payloads.Poll{
			ID:              uuid.New(),
			ContentableID:   newPost.ID,
			ContentableType: payloads.ContentablePollPost,
			Question:        *utils.MakeLocalizedString(defaultLanguage, "Pool Question"),
			Duration:        postForm.PollDuration,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		for _, choiceLabel := range postForm.PollOptions {
			fmt.Println("ANKET SECENEKLERI:", choiceLabel)
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

		newPost.Poll = poll
	}

	var locationPost *global_shared.Location // varsayılan olarak nil
	var locationPoint *extensions.PostGISPoint

	// location
	if len(postForm.LocationAddress) > 0 {

		if postForm.LocationLat != 0 && postForm.LocationLng != 0 {
			locationPoint = &extensions.PostGISPoint{
				Lat: postForm.LocationLat,
				Lng: postForm.LocationLng,
			}
		}

		locationPost = &global_shared.Location{
			ID:              uuid.New(),
			ContentableType: shared.LocationOwnerPost,
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
			ContentableType: shared.LocationOwnerEvent,
			ContentableID:   newPost.ID,
			Address:         &postForm.LocationAddress,
			Latitude:        &postForm.LocationLat,
			Longitude:       &postForm.LocationLng,
			LocationPoint:   locationPoint,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		evt.Location = locationEvent
		evt.LocationID = &locationEvent.ID
		if err := tx.Save(evt).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		newPost.Event = evt
	}

	newPost.Location = locationPost
	if err := tx.Save(newPost).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	fmt.Println("NEW POST", newPost.ID)
	return newPost, nil
}

func (s *PostService) GetPostByID(id uuid.UUID) (*post.Post, error) {
	postData, err := s.postRepo.GetPostByID(id)
	if err != nil {
		return nil, fmt.Errorf("GetPostByID error: %w", err)
	}
	return postData, nil
}
