package services

import (
	"bifrost/extensions"
	"bifrost/helpers"
	"bifrost/models/user"
	"bifrost/models/user/payloads"
	"bifrost/repositories"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Register işlemi
func (s *UserService) Register(form map[string][]string) (*user.User, string, error) {
	// Orientation
	orientationKey := form["body[orientation]"][0]
	var orientation payloads.SexualOrientation
	if err := s.repo.DB().Preload("Translations").
		Where("id = ?", orientationKey).First(&orientation).Error; err != nil {
		return nil, "", errors.New("invalid sexual orientation")
	}

	// BirthDate
	dateOfBirth, err := time.Parse("2006-01-02", form["body[birthDate]"][0])
	if err != nil {
		return nil, "", errors.New("invalid birthDate")
	}

	// Location
	lat, err := strconv.ParseFloat(form["body[location][lat]"][0], 64)
	if err != nil {
		return nil, "", errors.New("invalid latitude")
	}
	lng, err := strconv.ParseFloat(form["body[location][lng]"][0], 64)
	if err != nil {
		return nil, "", errors.New("invalid longitude")
	}

	node, err := helpers.NewNode(1)
	if err != nil {
		return nil, "", fmt.Errorf("failed to create snowflake node: %w", err)
	}

	userObj := &user.User{
		PublicID:            node.Generate().Int64(),
		UserName:            form["body[name]"][0],
		DisplayName:         form["body[nickname]"][0],
		DateOfBirth:         &dateOfBirth,
		SexualOrientationID: &orientation.ID,
		SexualOrientation:   &orientation,
		Location: &user.LocationData{
			CountryCode: form["body[location][country_code]"][0],
			CountryName: form["body[location][country_name]"][0],
			City:        form["body[location][city]"][0],
			Region:      form["body[location][region]"][0],
			Lat:         lat,
			Lng:         lng,
			Display:     form["body[location][display]"][0],
			Timezone:    form["body[location][timezone]"][0],
		},
		LocationPoint: &extensions.PostGISPoint{
			Lat: lat,
			Lng: lng,
		},
	}

	if err := s.repo.Create(userObj); err != nil {
		return nil, "", err
	}

	fmt.Println("INSERT:FANTASIES")

	// FANTASIES ekleme
	if len(form) > 0 {
		var fantasyIDs []uuid.UUID
		for k, v := range form {
			if strings.HasPrefix(k, "body[fantasies]") {
				for _, val := range v {
					fantasyUUID, err := uuid.Parse(val)
					if err != nil {
						fmt.Println("Skipping invalid fantasy ID:", val)
						continue
					}
					fantasyIDs = append(fantasyIDs, fantasyUUID)
				}
			}

		}

		fmt.Println(fantasyIDs)

		if len(fantasyIDs) > 0 {
			userFantasies := make([]*payloads.UserFantasy, len(fantasyIDs))
			fmt.Println(fantasyIDs)
			for i, fID := range fantasyIDs {
				userFantasies[i] = &payloads.UserFantasy{
					UserID:    userObj.ID,
					FantasyID: fID,
				}
			}

			// Toplu insert
			if err := s.repo.DB().Create(&userFantasies).Error; err != nil {
				return nil, "", fmt.Errorf("failed to insert user fantasies: %w", err)
			}
		}
	}

	userInfo, err := s.GetUserByID(userObj.ID)
	if err != nil {
		return nil, "", err
	}
	token, err := helpers.GenerateUserJWT(userObj.ID, userObj.Email)
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
