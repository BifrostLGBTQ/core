package user

import (
	"bifrost/constants"
	"bifrost/models/media"
	"bifrost/models/shared"
	"bifrost/models/user/payloads"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

const (
	FollowStatusFollowing constants.FollowStatus = "following"
	FollowStatusBlocked   constants.FollowStatus = "blocked"
	FollowStatusMuted     constants.FollowStatus = "muted"
	FollowStatusRequested constants.FollowStatus = "requested"
)

// === FOLLOW ===
type Follow struct {
	ID         uuid.UUID              `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	FollowerID uuid.UUID              `gorm:"type:uuid;index;not null" json:"follower_id"`
	FolloweeID uuid.UUID              `gorm:"type:uuid;index;not null" json:"followee_id"`
	Status     constants.FollowStatus `gorm:"type:varchar(20);default:'following';index" json:"status"`

	Follower *User `gorm:"foreignKey:FollowerID" json:"follower,omitempty"`
	Followee *User `gorm:"foreignKey:FolloweeID" json:"followee,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// === LIKE ===
type Like struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	LikerID uuid.UUID `gorm:"type:uuid;index;not null" json:"liker_id"`
	LikedID uuid.UUID `gorm:"type:uuid;index;not null" json:"liked_id"`
	IsMatch bool      `gorm:"default:false" json:"is_match"`

	Liker *User `gorm:"foreignKey:LikerID" json:"liker,omitempty"`
	Liked *User `gorm:"foreignKey:LikedID" json:"liked,omitempty"`

	CreatedAt time.Time `json:"created_at"`
}

// === FAVORITE ===
type Favorite struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID     uuid.UUID `gorm:"type:uuid;index;not null" json:"user_id"`
	FavoriteID uuid.UUID `gorm:"type:uuid;index;not null" json:"favorite_id"`
	Note       *string   `gorm:"type:text" json:"note,omitempty"`

	User     *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Favorite *User `gorm:"foreignKey:FavoriteID" json:"favorite,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// === MATCH ===
type Match struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserAID uuid.UUID `gorm:"type:uuid;index;not null" json:"user_a_id"`
	UserBID uuid.UUID `gorm:"type:uuid;index;not null" json:"user_b_id"`

	UserA *User `gorm:"foreignKey:UserAID" json:"user_a,omitempty"`
	UserB *User `gorm:"foreignKey:UserBID" json:"user_b,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Block struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	BlockerID uuid.UUID `gorm:"type:uuid;index;not null" json:"blocker_id"` // Engelleyen kullanıcı
	BlockedID uuid.UUID `gorm:"type:uuid;index;not null" json:"blocked_id"` // Engellenen kullanıcı
	Reason    *string   `gorm:"type:text" json:"reason,omitempty"`          // Opsiyonel açıklama

	Blocker *User `gorm:"foreignKey:BlockerID" json:"blocker,omitempty"`
	Blocked *User `gorm:"foreignKey:BlockedID" json:"blocked,omitempty"`

	CreatedAt time.Time `json:"created_at"`
}

type SocialRelations struct {
	Follows        []*Follow   `json:"-" gorm:"foreignKey:FollowerID"`
	Followers      []*Follow   `json:"-" gorm:"foreignKey:FolloweeID"`
	Likes          []*Like     `json:"-" gorm:"foreignKey:LikerID"`
	LikedBy        []*Like     `json:"-" gorm:"foreignKey:LikedID"`
	Matches        []*Match    `json:"-" gorm:"foreignKey:UserAID"`
	Favorites      []*Favorite `json:"-" gorm:"foreignKey:UserID"`
	FavoritedBy    []*Favorite `json:"-" gorm:"foreignKey:FavoriteID"`
	BlockedUsers   []*Block    `gorm:"foreignKey:BlockerID" json:"blocked_users,omitempty"`
	BlockedByUsers []*Block    `gorm:"foreignKey:BlockedID" json:"blocked_by_users,omitempty"`
}

type TravelData struct {
	VisitedCountries pq.StringArray `gorm:"type:text[]" json:"visited_countries"`
	TravelFrequency  string         `json:"travel_frequency"`                   // örn: "aylık"
	FavoritePlaces   pq.StringArray `gorm:"type:text[]" json:"favorite_places"` // opsiyonel
}

// Ziyaret Edilen Ülkeler
type CountryVisit struct {
	ISOCode   string    `json:"iso_code"`             // Örn: "FR"
	Name      string    `json:"name"`                 // Örn: "France"
	VisitedAt time.Time `json:"visited_at,omitempty"` // İsteğe bağlı
	Notes     string    `json:"notes,omitempty"`
}

// Sevilen Şehirler
type FavoriteCity struct {
	City      string    `json:"city"`                 // Örn: "Tokyo"
	Country   string    `json:"country"`              // Örn: "Japan"
	ISOCode   string    `json:"iso_code"`             // Örn: "JP"
	Reason    string    `json:"reason,omitempty"`     // Neden favori?
	LastVisit time.Time `json:"last_visit,omitempty"` // Son ziyaret tarihi
}

// Seyahat Planı
type TravelPlan struct {
	City        string                  `json:"city"`     // Örn: "Barcelona"
	Country     string                  `json:"country"`  // Örn: "Spain"
	ISOCode     string                  `json:"iso_code"` // Örn: "ES"
	StartDate   time.Time               `json:"start_date"`
	EndDate     time.Time               `json:"end_date"`
	Purpose     constants.TravelPurpose `json:"purpose,omitempty"` // Enum: vacation, work, etc.
	WithFriends bool                    `json:"with_friends"`      // Yalnız mı gidiyor?
	Notes       string                  `json:"notes,omitempty"`
	IsPublic    bool                    `json:"is_public"` // Profilde gözükebilir mi?
}

type User struct {
	ID                  uuid.UUID                    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	PublicID            int64                        `gorm:"uniqueIndex;not null" json:"public_id"`
	SocketID            *string                      `json:"socket_id,omitempty"`
	UserName            string                       `json:"username"`
	DisplayName         string                       `json:"displayname"`
	Email               string                       `json:"email"`
	Password            string                       `json:"-"` // gizli tutulmalı
	ProfileImageURL     *string                      `json:"profile_image_url,omitempty"`
	Bio                 *string                      `json:"bio,omitempty"`
	DateOfBirth         *time.Time                   `json:"date_of_birth,omitempty"`
	Gender              constants.GenderIdentity     `json:"gender"`
	SexualOrientation   *payloads.SexualOrientation  `gorm:"foreignKey:SexualOrientationID" json:"sexual_orientation"`
	SexualOrientationID *uuid.UUID                   `gorm:"type:uuid;index" json:"-"`
	RoleInSex           constants.SexRole            `json:"sex_role"`
	RelationshipStatus  constants.RelationshipStatus `json:"relationship_status"`
	UserRole            constants.UserRole           `json:"user_role"`
	IsActive            bool                         `json:"is_active"`
	CreatedAt           time.Time                    `json:"created_at"`
	UpdatedAt           time.Time                    `json:"updated_at"`
	LastOnline          *time.Time                   `json:"last_online,omitempty"`
	Location            *shared.Location             `gorm:"polymorphic:Contentable;polymorphicValue:user;constraint:OnDelete:CASCADE" json:"location,omitempty"`

	DefaultLanguage string `gorm:"type:varchar(8);default:'en'" json:"default_language"`

	AvatarID *uuid.UUID   `json:"avatar_id,omitempty"`
	Avatar   *media.Media `gorm:"constraint:OnDelete:SET NULL;foreignKey:AvatarID;references:ID" json:"avatar,omitempty"`

	CoverID *uuid.UUID   `json:"cover_id,omitempty"`
	Cover   *media.Media `gorm:"constraint:OnDelete:SET NULL;foreignKey:CoverID;references:ID" json:"cover,omitempty"`

	// BDSM
	BDSMInterest constants.BDSMInterest `json:"bdsm_interest,omitempty"`
	BDSMRole     constants.BDSMRole     `json:"bdsm_role,omitempty"`

	// Alkol ve Sigara kullanımı
	Smoking  constants.SmokingHabit  `json:"smoking,omitempty"`
	Drinking constants.DrinkingHabit `json:"drinking,omitempty"`

	// Hobi ve Eğlence alanları (liste şeklinde)

	Languages     pq.StringArray           `gorm:"type:text[]" json:"languages"`
	Hobbies       pq.StringArray           `gorm:"type:text[]" json:"hobbies,omitempty"`
	MoviesGenres  pq.StringArray           `gorm:"type:text[]" json:"movies_genres,omitempty"`
	TVShowsGenres pq.StringArray           `gorm:"type:text[]" json:"tv_shows_genres,omitempty"`
	TheaterGenres pq.StringArray           `gorm:"type:text[]" json:"theater_genres,omitempty"`
	CinemaGenres  pq.StringArray           `gorm:"type:text[]" json:"cinema_genres,omitempty"`
	ArtInterests  pq.StringArray           `gorm:"type:text[]" json:"art_interests,omitempty"`
	Entertainment pq.StringArray           `gorm:"type:text[]" json:"entertainment,omitempty"`
	Fantasies     []*payloads.UserFantasy  `gorm:"foreignKey:UserID" json:"fantasies,omitempty"`
	Interests     []*payloads.UserInterest `gorm:"foreignKey:UserID" json:"interests,omitempty"`

	Travel TravelData `gorm:"embedded;embeddedPrefix:travel_" json:"travel"`
	//  Sosyal İlişkiler
	SocialRelations SocialRelations `json:"social,omitempty" gorm:"embedded;embeddedPrefix:social_"`
	Media           []*media.Media  `gorm:"polymorphic:Owner;polymorphicValue:user;constraint:OnDelete:CASCADE" json:"media,omitempty"`
	DeletedAt       gorm.DeletedAt  `gorm:"index" json:"deleted_at,omitempty"`
	jwt.StandardClaims
}

func (User) TableName() string {
	return "users"
}

func (TravelPlan) TableName() string {
	return "user_travel_plans"
}

func (FavoriteCity) TableName() string {
	return "user_favorite_cities"
}

func (Block) TableName() string {
	return "user_blocks"
}

func (Match) TableName() string {
	return "user_matches"
}

func (Like) TableName() string {
	return "user_likes"
}

func (Follow) TableName() string {
	return "user_follows"
}

func (CountryVisit) TableName() string {
	return "user_country_visits"
}

func (Favorite) TableName() string {
	return "user_favorites"
}
