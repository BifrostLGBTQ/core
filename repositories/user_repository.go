package repositories

import (
	"bifrost/extensions"
	"bifrost/models/user"

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
		LocationPoint: extensions.PostGISPoint{
			Lng: 83.96632795978059,
			Lat: 28.2052577611216,
		},
	}

	return r.db.Create(&user).Error
}

func (r *UserRepository) Create(user *user.User) error {
	return r.db.Create(user).Error
}

// ID ile kullanıcıyı al
func (r *UserRepository) GetByID(userID uuid.UUID) (*user.User, error) {
	var u user.User
	err :=
		r.db.
			Preload("Fantasies.Fantasy.Translations").
			Preload("Media").
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
