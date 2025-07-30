package payloads

import (
	"bifrost/models"
	"time"

	"github.com/google/uuid"
)

// --- Poll (Anketler için) ---

type Poll struct {
	ID                    uuid.UUID    `gorm:"type:uuid;primaryKey"`
	Question              string       `gorm:"size:512;not null"`
	CreatorID             uuid.UUID    `gorm:"type:uuid;not null;index"`
	Creator               models.User  `gorm:"foreignKey:CreatorID"`
	IsAnonymous           bool         `gorm:"default:false"` // Oy veren kullanıcılar gizli mi?
	AllowsMultipleAnswers bool         `gorm:"default:false"` // Çoklu oy destekli mi?
	Options               []PollOption `gorm:"foreignKey:PollID"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
type PollOption struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey"`
	PollID    uuid.UUID  `gorm:"type:uuid;not null;index"`
	Poll      Poll       `gorm:"foreignKey:PollID"`
	Text      string     `gorm:"size:256;not null"`
	VoteCount int        `gorm:"default:0"`           // Kolay erişim için cache'lenmiş sayım
	Votes     []PollVote `gorm:"foreignKey:OptionID"` // Kim oy verdi
}

// --- Anket oyları (katılımcılar) ---

type PollVote struct {
	ID       uuid.UUID  `gorm:"type:uuid;primaryKey"`
	OptionID uuid.UUID  `gorm:"type:uuid;not null;index"`
	Option   PollOption `gorm:"foreignKey:OptionID"`

	UserID uuid.UUID   `gorm:"type:uuid;not null;index"`
	User   models.User `gorm:"foreignKey:UserID"`

	VotedAt time.Time `gorm:"autoCreateTime"`
}

func (Poll) TableName() string {
	return "messages_pools"
}

func (PollVote) TableName() string {
	return "messages_pool_votes"
}

func (PollOption) TableName() string {
	return "messages_pool_options"
}
