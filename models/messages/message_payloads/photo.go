package message_payloads

import (
	"time"

	"github.com/google/uuid"
)

type Photo struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	URL       string    `gorm:"size:512;not null"`
	Width     int       `gorm:"not null"`
	Height    int       `gorm:"not null"`
	CreatedAt time.Time
}

func (Photo) TableName() string {
	return "messages_photo"
}
