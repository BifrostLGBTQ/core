package services

import (
	"bifrost/constants"
	"bifrost/extensions"
	"bifrost/helpers"
	"bifrost/models/media"
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
	repo      *repositories.UserRepository
	mediaRepo *repositories.MediaRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Register işlemi
func (s *UserService) Register(request map[string][]string) (*user.User, string, error) {

	type RegisterForm struct {
		Name      string `form:"name"`
		Nickname  string `form:"nickname"`
		Password  string `form:"password"`
		BirthDate string `form:"birthDate"` // string veya time.Time

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

	// BirthDate
	dateOfBirth, err := time.Parse("2006-01-02", formData.BirthDate)
	if err != nil {
		return nil, "", errors.New("invalid birthDate")
	}

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

		ID:          UserID,
		PublicID:    node.Generate().Int64(),
		UserName:    formData.Name,
		DisplayName: formData.Nickname,
		DateOfBirth: &dateOfBirth,
	}

	if err := s.repo.Create(userObj); err != nil {
		return nil, "", err
	}

	fmt.Println("INSERT:FANTASIES")

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

func (s *UserService) UpdateAvatar(file *multipart.FileHeader, user *user.User) (*media.Media, error) {
	newMedia, err := s.mediaRepo.AddMedia(
		s.repo.DB(),
		user.ID,
		media.OwnerUser,
		media.RoleAvatar,
		file,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to upload avatar: %w", err)
	}

	// User tablosunu güncelle
	user.AvatarID = &newMedia.ID
	user.Avatar = newMedia

	if err := s.repo.UpdateUser(user); err != nil {
		return nil, fmt.Errorf("failed to update user avatar: %w", err)
	}
	return newMedia, nil
}

func (s *UserService) UpdateCover(file *multipart.FileHeader, user *user.User) (*media.Media, error) {
	//
	newMedia, err := s.mediaRepo.AddMedia(
		s.repo.DB(),
		user.ID,
		media.OwnerUser,
		media.RoleAvatar,
		file,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to upload avatar: %w", err)
	}
	user.CoverID = &newMedia.ID
	user.Cover = newMedia

	if err := s.repo.UpdateUser(user); err != nil {
		return nil, fmt.Errorf("failed to update user avatar: %w", err)
	}
	return newMedia, nil
}

func (s *UserService) AddStory(file *multipart.FileHeader, user *user.User) (*payloads.Story, error) {
	//
	fmt.Println("Upload COVER")
	return nil, nil
}

func (s *UserService) GetAttribute(attributeID uuid.UUID) (*payloads.Attribute, error) {
	return s.repo.GetAttribute(attributeID)
}

func (s *UserService) GetInterestItem(interestId uuid.UUID) (*payloads.InterestItem, error) {
	return s.repo.GetInterestItem(interestId)
}

// Kullanıcı ID ile getir
func (s *UserService) GetFantasy(id uuid.UUID) (*payloads.Fantasy, error) {
	return s.repo.GetFantasy(id)
}

func (s *UserService) UpsertUserSexualIdentify(
	userID uuid.UUID,
	genderIDs []string,
	sexualIDs []string,
	sexRoleIDs []string,
) error {

	// Kullanıcıyı repo'dan çekiyoruz (ilişkilerle birlikte)
	user, err := s.repo.GetUserWithSexualRelations(userID)
	if err != nil {
		return err
	}

	// GenderIdentities güncelle
	if genderIDs != nil {
		if len(genderIDs) == 0 {
			if err := s.repo.ClearGenderIdentities(user); err != nil {
				return err
			}
		} else {
			ids, err := parseUUIDs(genderIDs)
			if err != nil {
				return err
			}
			if err := s.repo.ReplaceGenderIdentities(user, ids); err != nil {
				return err
			}
		}
	}

	// SexualOrientations güncelle
	if sexualIDs != nil {
		if len(sexualIDs) == 0 {
			if err := s.repo.ClearSexualOrientations(user); err != nil {
				return err
			}
		} else {
			ids, err := parseUUIDs(sexualIDs)
			if err != nil {
				return err
			}
			if err := s.repo.ReplaceSexualOrientations(user, ids); err != nil {
				return err
			}
		}
	}

	// SexRole güncelle (tek ilişki)
	if sexRoleIDs != nil {
		if len(sexRoleIDs) == 0 {
			if err := s.repo.ClearSexRole(user); err != nil {
				return err
			}
		} else {
			id, err := uuid.Parse(sexRoleIDs[0])
			if err != nil {
				return err
			}

			fmt.Println("SET ROLE SEX GHERE", user.DisplayName, id)
			if err := s.repo.SetSexRole(user, id); err != nil {
				fmt.Println("SET ROLE HATA OLDU GHERE", user.DisplayName)

				return err
			}
		}
	}

	return nil
}

func parseUUIDs(strIDs []string) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	for _, strID := range strIDs {
		id, err := uuid.Parse(strID)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (s *UserService) UpsertUserAttribute(attr *payloads.UserAttribute) error {
	if attr == nil {
		return fmt.Errorf("attribute cannot be nil")
	}

	if attr.UserID == uuid.Nil {
		return fmt.Errorf("user_id is required")
	}

	if attr.AttributeID == uuid.Nil {
		return fmt.Errorf("attribute_id is required")
	}

	// Repository'yi çağır
	err := s.repo.UpsertUserAttribute(attr)
	if err != nil {
		return fmt.Errorf("failed to upsert user attribute: %w", err)
	}

	return nil
}

func (s *UserService) UpsertUserInterest(interest *payloads.UserInterest) error {
	if interest == nil {
		return fmt.Errorf("attribute cannot be nil")
	}

	if interest.UserID == uuid.Nil {
		return fmt.Errorf("user_id is required")
	}

	if interest.InterestItemID == uuid.Nil {
		return fmt.Errorf("attribute_id is required")
	}

	// Repository'yi çağır
	err := s.repo.ToggleUserInterest(interest)
	if err != nil {
		return fmt.Errorf("failed to upsert user attribute: %w", err)
	}

	return nil
}

func (s *UserService) UpsertUserFantasy(fantasy *payloads.UserFantasy) error {
	if fantasy == nil {
		return fmt.Errorf("fantasy cannot be nil")
	}

	if fantasy.UserID == uuid.Nil {
		return fmt.Errorf("user_id is required")
	}

	if fantasy.FantasyID == uuid.Nil {
		return fmt.Errorf("fantasy is required")
	}

	// Repository'yi çağır
	err := s.repo.ToggleUserFantasy(fantasy)
	if err != nil {
		return fmt.Errorf("failed to upsert user attribute: %w", err)
	}

	return nil
}
