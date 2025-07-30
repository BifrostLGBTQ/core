package payloads

import (
	"time"

	"github.com/google/uuid"
)

type Gift struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ReceiverID uuid.UUID `gorm:"type:uuid;not null;index"` // Hediye alan kişi
	Name       string    `gorm:"size:128;not null"`
	Amount     int       `gorm:"not null"`
	Note       *string   `gorm:"size:256"`
	CreatedAt  time.Time
}
