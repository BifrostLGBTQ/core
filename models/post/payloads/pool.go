package payloads

import (
	"bifrost/models/post/shared"
	"time"

	"github.com/google/uuid"
)

const (
	ContentablePollPost = "post"
)

type Poll struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	PostID          uuid.UUID `gorm:"type:uuid;index;not null"` // Hangi posta ait
	ContentableID   uuid.UUID
	ContentableType string
	Question        shared.LocalizedString `gorm:"type:jsonb"`
	Duration        int                    `gorm:"default:0"` // Poll süresi (dakika/isteğe bağlı)
	CreatedAt       time.Time
	UpdatedAt       time.Time

	Choices []PollChoice `gorm:"foreignKey:PollID;constraint:OnDelete:CASCADE"`
}

type PollChoice struct {
	ID        uuid.UUID              `gorm:"type:uuid;primaryKey"`
	PollID    uuid.UUID              `gorm:"type:uuid;index;not null"`
	Label     shared.LocalizedString `gorm:"type:jsonb"` // {"en":"Yes","tr":"Evet"}
	VoteCount int                    `gorm:"default:0"`

	Votes []PollVote `gorm:"foreignKey:ChoiceID;constraint:OnDelete:CASCADE"`
}

type PollVote struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	ChoiceID  uuid.UUID `gorm:"type:uuid;index;not null"`
	UserID    uuid.UUID `gorm:"type:uuid;index;not null"` // Oy veren kullanıcı
	CreatedAt time.Time
}
