package services

import (
	"bifrost/constants"
	"bifrost/extensions"
	"bifrost/helpers"
	global_shared "bifrost/models/shared"
	"bifrost/models/user"
	"bifrost/models/user/payloads"
	"bifrost/repositories"
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	form "github.com/go-playground/form/v4"
	"github.com/google/uuid"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Register işlemi
func (s *UserService) Register(request map[string][]string) (*user.User, string, error) {

	type RegisterForm struct {
		Name        string `form:"name"`
		Nickname    string `form:"nickname"`
		Password    string `form:"password"`
		BirthDate   string `form:"birthDate"` // string veya time.Time
		Orientation string `form:"orientation"`

		// Fantasies array
		Fantasies []string `form:"fantasies[]"`

		// Nested location
		CountryCode string  `form:"location[country_code]"`
		Country     string  `form:"location[country_name]"`
		City        string  `form:"location[city]"`
		Region      string  `form:"location[region]"`
		Lat         float64 `form:"location[lat]"`
		Lng         float64 `form:"location[lng]"`
		Timezone    string  `form:"location[timezone]"`
		Display     string  `form:"location[display]"`
		Address     string  `form:"location[address]"` // varsa
	}
	decoder := form.NewDecoder()
	var formData RegisterForm

	// formValues map[string][]string şeklinde gelecek
	if err := decoder.Decode(&formData, request); err != nil {
		return nil, "", err
	}

	fmt.Println("DATA", formData)

	fmt.Println("Orientation:BEGIN")
	// Orientation
	orientationKey := formData.Orientation
	var orientation payloads.SexualOrientation
	if err := s.repo.DB().Preload("Translations").
		Where("id = ?", orientationKey).First(&orientation).Error; err != nil {
		return nil, "", errors.New("invalid sexual orientation")
	}
	fmt.Println("Orientation:END")

	// BirthDate
	dateOfBirth, err := time.Parse("2006-01-02", formData.BirthDate)
	if err != nil {
		return nil, "", errors.New("invalid birthDate")
	}

	fmt.Println("ORIENTATION", orientationKey, "BITYH", dateOfBirth)

	node, err := helpers.NewNode(1)
	if err != nil {
		return nil, "", fmt.Errorf("failed to create snowflake node: %w", err)
	}

	locationPoint := &extensions.PostGISPoint{
		Lat: formData.Lat,
		Lng: formData.Lng,
	}

	UserID := uuid.New()
	locationUser := &global_shared.Location{
		ID:              uuid.New(),
		ContentableType: global_shared.LocationOwnerUser,
		ContentableID:   UserID,

		CountryCode:   &formData.CountryCode,
		Country:       &formData.Country,
		City:          &formData.City,
		Region:        &formData.Region,
		Display:       &formData.Display,
		Timezone:      &formData.Timezone,
		Address:       &formData.Address,
		Latitude:      &formData.Lat,
		Longitude:     &formData.Lng,
		LocationPoint: locationPoint,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := s.repo.UpsertLocation(locationUser); err != nil {
		return nil, "", err
	}

	userObj := &user.User{

		ID:                  UserID,
		PublicID:            node.Generate().Int64(),
		UserName:            formData.Name,
		DisplayName:         formData.Nickname,
		DateOfBirth:         &dateOfBirth,
		SexualOrientationID: &orientation.ID,
		SexualOrientation:   &orientation,
	}

	if err := s.repo.Create(userObj); err != nil {
		return nil, "", err
	}

	fmt.Println("INSERT:FANTASIES")

	if len(formData.Fantasies) > 0 {
		userFantasies := make([]*payloads.UserFantasy, len(formData.Fantasies))
		for i, fID := range formData.Fantasies {

			fantasyUUID, err := uuid.Parse(fID)
			if err != nil {
				fmt.Println("Skipping invalid fantasy ID:", fID)
				continue
			}
			userFantasies[i] = &payloads.UserFantasy{
				UserID:    userObj.ID,
				FantasyID: fantasyUUID,
			}
		}

		// Toplu insert
		if err := s.repo.DB().Create(&userFantasies).Error; err != nil {
			return nil, "", fmt.Errorf("failed to insert user fantasies: %w", err)
		}
	}

	userInfo, err := s.GetUserByID(userObj.ID)
	if err != nil {
		return nil, "", err
	}
	token, err := helpers.GenerateUserJWT(userObj.ID, userObj.PublicID)
	if err != nil {
		return nil, "", err
	}

	return userInfo, token, nil
}

// Kullanıcı ID ile getir
func (s *UserService) GetUserByID(id uuid.UUID) (*user.User, error) {
	return s.repo.GetByID(id)
}

// Register işlemi
func (s *UserService) Test() {

	if err := s.repo.TestUser(); err != nil {
		return
	}

}

func (s *UserService) Follow(followerID, followeeID string) error {
	fID, err := uuid.Parse(followerID)
	if err != nil {
		return errors.New(constants.ErrInvalidInput.String())
	}
	feID, err := uuid.Parse(followeeID)
	if err != nil {
		return errors.New(constants.ErrInvalidInput.String())
	}

	if err := s.repo.Follow(fID, feID); err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (s *UserService) Unfollow(followerID, followeeID string) error {
	fID, err := uuid.Parse(followerID)
	if err != nil {
		return errors.New(constants.ErrInvalidInput.String())
	}
	feID, err := uuid.Parse(followeeID)
	if err != nil {
		return errors.New(constants.ErrInvalidInput.String())
	}

	if err := s.repo.Unfollow(fID, feID); err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (s *UserService) UpdateAvatar(file *multipart.FileHeader, author *user.User) (*user.User, error) {
	//
	fmt.Println("Upload AVATAR")
	return nil, nil
}

func (s *UserService) UpdateCover(file *multipart.FileHeader, author *user.User) (*user.User, error) {
	//
	fmt.Println("Upload COVER")
	return nil, nil
}

func (s *UserService) AddStory(file *multipart.FileHeader, author *user.User) (*user.User, error) {
	//
	fmt.Println("Upload COVER")
	return nil, nil
}
