package payloads

import (
	"time"

	"github.com/google/uuid"
)

type Audio struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	URL       string    `gorm:"size:512;not null"`
	Duration  int       `gorm:"not null"`
	CreatedAt time.Time
}
