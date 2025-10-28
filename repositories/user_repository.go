package repositories

import (
	"bifrost/constants"
	global_shared "bifrost/models/shared"
	"bifrost/models/user"
	"errors"
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
	user := user.User{
		UserName:    "testUser",
		DisplayName: "testUser",
	}

	return r.db.Create(&user).Error
}

func (r *UserRepository) Create(user *user.User) error {
	return r.db.Create(user).Error
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
	var existing user.Follow
	if err := r.db.
		Where("follower_id = ? AND followee_id = ?", followerID, followeeID).
		First(&existing).Error; err == nil {
		return errors.New(constants.ErrDuplicateResource.String()) // Zaten takip ediyor
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(constants.ErrDatabaseError.String()) // DB hatası
	}

	follow := user.Follow{
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
		Delete(&user.Follow{}).Error; err != nil {
		return errors.New(constants.ErrDatabaseError.String())
	}
	return nil
}

// ID ile kullanıcıyı al
func (r *UserRepository) GetByID(userID uuid.UUID) (*user.User, error) {
	var u user.User
	err :=
		r.db.
			Preload("Fantasies.Fantasy.Translations").
			Preload("Avatar").
			Preload("Cover").
			Preload("SexualOrientation").
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

func (r *UserRepository) GetUserByPublicId(userID int64) (*user.User, error) {
	var u user.User
	err :=
		r.db.
			Preload("Fantasies.Fantasy.Translations").
			Preload("Avatar").
			Preload("Cover").
			Preload("SexualOrientation").
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
