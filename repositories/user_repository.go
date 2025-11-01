package repositories

import (
	"bifrost/constants"
	global_shared "bifrost/models/shared"
	userModel "bifrost/models/user"
	"bifrost/models/user/payloads"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) DB() *gorm.DB {
	return r.db
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) TestUser() error {
	user := userModel.User{
		UserName:    "testUser",
		DisplayName: "testUser",
	}

	return r.db.Create(&user).Error
}

func (r *UserRepository) Create(user *userModel.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) UpdateUser(u *userModel.User) error {
	return r.db.Save(u).Error
}

func (r *UserRepository) DeleteUser(userID uuid.UUID) error {
	return r.db.
		Where("id = ?", userID).
		Delete(&userModel.User{}).Error
}

func (r *UserRepository) Login(username string, password string) error {
	return nil
}

func (r *UserRepository) LoginViaToken(token string) error {
	return nil
}

// Kullanıcıyı takip et
func (r *UserRepository) Follow(followerID, followeeID uuid.UUID) error {
	if followerID == followeeID {
		return errors.New(constants.ErrInvalidAction.String()) // Kendini takip edemezsin
	}

	// Zaten takip ediyor mu kontrol et
	var existing userModel.Follow
	if err := r.db.
		Where("follower_id = ? AND followee_id = ?", followerID, followeeID).
		First(&existing).Error; err == nil {
		return errors.New(constants.ErrDuplicateResource.String()) // Zaten takip ediyor
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(constants.ErrDatabaseError.String()) // DB hatası
	}

	follow := userModel.Follow{
		FollowerID: followerID,
		FolloweeID: followeeID,
		Status:     "following",
	}

	if err := r.db.Create(&follow).Error; err != nil {
		return errors.New(constants.ErrDatabaseError.String())
	}

	return nil
}

// Takipten çık
func (r *UserRepository) Unfollow(followerID, followeeID uuid.UUID) error {
	if err := r.db.
		Where("follower_id = ? AND followee_id = ?", followerID, followeeID).
		Delete(&userModel.Follow{}).Error; err != nil {
		return errors.New(constants.ErrDatabaseError.String())
	}
	return nil
}

// ID ile kullanıcıyı al
func (r *UserRepository) GetByID(userID uuid.UUID) (*userModel.User, error) {
	var u userModel.User

	err :=
		r.db.
			Preload("Fantasies.Fantasy.Translations").
			Preload("Interests.InterestItem.Interest").
			Preload("Avatar.File").
			Preload("Cover.File").
			Preload("GenderIdentities").
			Preload("SexualOrientations").
			Preload("SexualRole").
			Preload("UserAttributes.Attribute").
			Preload("Media").
			Preload("SocialRelations.Follows").
			Preload("SocialRelations.Followers").
			Preload("SocialRelations.Likes").
			Preload("SocialRelations.LikedBy").
			Preload("SocialRelations.Matches").
			Preload("SocialRelations.Favorites").
			Preload("SocialRelations.FavoritedBy").
			Preload("SocialRelations.BlockedUsers").
			Preload("SocialRelations.BlockedByUsers").
			First(&u, "id = ?", userID).Error

	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) GetUserByPublicId(userID int64) (*userModel.User, error) {
	var u userModel.User
	err :=
		r.db.
			Preload("Fantasies.Fantasy.Translations").
			Preload("Avatar").
			Preload("Cover").
			Preload("Media").
			Preload("SocialRelations.Follows").
			Preload("SocialRelations.Followers").
			Preload("SocialRelations.Likes").
			Preload("SocialRelations.LikedBy").
			Preload("SocialRelations.Matches").
			Preload("SocialRelations.Favorites").
			Preload("SocialRelations.FavoritedBy").
			Preload("SocialRelations.BlockedUsers").
			Preload("SocialRelations.BlockedByUsers").
			First(&u, "public_id = ?", userID).Error

	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) UpsertLocation(location *global_shared.Location) error {
	if location.ID == uuid.Nil {
		location.ID = uuid.New()
	}

	location.UpdatedAt = time.Now()
	if location.CreatedAt.IsZero() {
		location.CreatedAt = time.Now()
	}

	// Polymorphic owner_type + owner_id eşleşmesini kontrol et
	var existing global_shared.Location
	err := r.db.Where("contentable_type = ? AND contentable_id = ?", location.ContentableType, location.ContentableID).First(&existing).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Yeni ekle
			return r.db.Create(location).Error
		}
		return err
	}

	// Güncelle
	location.ID = existing.ID
	return r.db.Model(&existing).Updates(location).Error
}

func (r *UserRepository) CreateStory(st *payloads.Story) error {
	return r.db.Create(st).Error
}

func (r *UserRepository) GetUserStories(userID uuid.UUID) ([]*payloads.Story, error) {
	var stories []*payloads.Story
	if err := r.db.Preload("Media").Where("user_id = ? AND is_expired = false", userID).Order("created_at DESC").Find(&stories).Error; err != nil {
		return nil, err
	}
	return stories, nil
}

func (r *UserRepository) GetAllUserStories() ([]*payloads.Story, error) {
	var stories []*payloads.Story
	if err := r.db.Preload("Media").Where("is_expired = false").Order("created_at DESC").Find(&stories).Error; err != nil {
		return nil, err
	}
	return stories, nil
}

func (r *UserRepository) ExpireOldStories() error {
	return r.db.Model(&payloads.Story{}).
		Where("expires_at <= ? AND is_expired = false", gorm.Expr("NOW()")).
		Update("is_expired", true).Error
}

func (r *UserRepository) GetAttribute(attributeID uuid.UUID) (*payloads.Attribute, error) {
	var attr payloads.Attribute
	if err := r.db.Where("id = ?", attributeID).First(&attr).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &attr, nil
}

func (r *UserRepository) GetInterestItem(interestId uuid.UUID) (*payloads.InterestItem, error) {
	var interest payloads.InterestItem
	if err := r.db.Where("id = ?", interestId).First(&interest).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &interest, nil
}

func (r *UserRepository) GetFantasy(fantasyId uuid.UUID) (*payloads.Fantasy, error) {
	var fantasy payloads.Fantasy
	if err := r.db.Where("id = ?", fantasyId).First(&fantasy).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &fantasy, nil
}

func (r *UserRepository) UpsertUserAttribute(attr *payloads.UserAttribute) error {

	fmt.Println("USER", attr.AttributeID, attr.UserID)
	attr.ID = uuid.New()
	if attr.AttributeID == uuid.Nil {
		return fmt.Errorf("invalid attribute")

	}

	if attr.UserID == uuid.Nil {
		return fmt.Errorf("invalid user")
	}

	var existing payloads.UserAttribute
	err := r.db.Where("user_id = ? AND category_type = ?", attr.UserID, attr.CategoryType).First(&existing).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Kayıt yoksa ekle
			if attr.ID == uuid.Nil {
				attr.ID = uuid.New()
			}
			return r.db.Create(attr).Error
		}
		return err
	}
	existing.AttributeID = attr.AttributeID
	existing.Notes = attr.Notes
	return r.db.Save(&existing).Error
}

func (r *UserRepository) ToggleUserInterest(interest *payloads.UserInterest) error {
	if interest.InterestItemID == uuid.Nil {
		return fmt.Errorf("invalid interest_item_id")
	}

	if interest.UserID == uuid.Nil {
		return fmt.Errorf("invalid user_id")
	}

	var existing payloads.UserInterest
	err := r.db.Where("user_id = ? AND interest_item_id = ?", interest.UserID, interest.InterestItemID).First(&existing).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Kayıt yok → ekle
		if interest.ID == uuid.Nil {
			interest.ID = uuid.New()
		}
		return r.db.Create(interest).Error
	} else if err != nil {
		return err
	}

	// Kayıt varsa → sil
	return r.db.Delete(&existing).Error
}

func (r *UserRepository) ToggleUserFantasy(fantasy *payloads.UserFantasy) error {
	if fantasy.FantasyID == uuid.Nil {
		return fmt.Errorf("invalid fantasy_id")
	}

	if fantasy.UserID == uuid.Nil {
		return fmt.Errorf("invalid user_id")
	}

	var existing payloads.UserFantasy
	err := r.db.Where("user_id = ? AND fantasy_id = ?", fantasy.UserID, fantasy.FantasyID).First(&existing).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Kayıt yok → ekle
		if fantasy.ID == uuid.Nil {
			fantasy.ID = uuid.New()
		}
		return r.db.Create(fantasy).Error
	} else if err != nil {
		return err
	}

	// Kayıt varsa → sil
	return r.db.Delete(&existing).Error
}

func (r *UserRepository) GetUserWithSexualRelations(userID uuid.UUID) (*userModel.User, error) {
	var user userModel.User
	err := r.db.Preload("GenderIdentities").
		Preload("SexualOrientations").
		Preload("SexualRole").
		First(&user, "id = ?", userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) ClearGenderIdentities(user *userModel.User) error {
	return r.db.Model(user).Association("GenderIdentities").Clear()
}

func (r *UserRepository) ReplaceGenderIdentities(user *userModel.User, ids []uuid.UUID) error {
	var genders []payloads.GenderIdentity
	for _, id := range ids {
		genders = append(genders, payloads.GenderIdentity{ID: id})
	}
	return r.db.Model(user).Association("GenderIdentities").Replace(genders)
}

func (r *UserRepository) ClearSexualOrientations(user *userModel.User) error {
	return r.db.Model(user).Association("SexualOrientations").Clear()
}

func (r *UserRepository) ReplaceSexualOrientations(user *userModel.User, ids []uuid.UUID) error {
	var sexuals []payloads.SexualOrientation
	for _, id := range ids {
		sexuals = append(sexuals, payloads.SexualOrientation{ID: id})
	}
	return r.db.Model(user).Association("SexualOrientations").Replace(sexuals)
}

func (r *UserRepository) ClearSexRole(user *userModel.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) SetSexRole(user *userModel.User, sexRoleID uuid.UUID) error {
	var dbUser userModel.User
	if err := r.db.First(&dbUser, "id = ?", user.ID).Error; err != nil {
		return err
	}
	dbUser.SexualRoleID = &sexRoleID
	if err := r.db.Save(&dbUser).Error; err != nil {
		return err
	}
	return nil
}
