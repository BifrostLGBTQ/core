package payloads

import (
	"bifrost/models"
	"time"
)

// --- Poll (Anketler için) ---

type Poll struct {
	ID        uint         `gorm:"primaryKey"`
	Question  string       `gorm:"size:512;not null"`
	Options   []PollOption `gorm:"foreignKey:PollID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// Anket tipi, gizli oy, çoklu seçim gibi özellikler eklenecekse buraya eklenir.
}

// --- Poll seçenekleri ---

type PollOption struct {
	ID        uint       `gorm:"primaryKey"`
	PollID    uint       `gorm:"index;not null"`
	Text      string     `gorm:"size:256;not null"`
	VoteCount uint       `gorm:"default:0"`
	Votes     []PollVote `gorm:"foreignKey:PollOptionID"`
}

// --- Anket oyları (katılımcılar) ---

type PollVote struct {
	ID           uint `gorm:"primaryKey"`
	PollOptionID uint `gorm:"index;not null"`
	PollOption   PollOption
	UserID       uint `gorm:"index;not null"`
	User         models.User
	VotedAt      time.Time `gorm:"autoCreateTime"`
}
