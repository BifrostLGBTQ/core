package message_payloads

import (
	"time"

	"github.com/google/uuid"
)

type Location struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Latitude  float64   `gorm:"type:decimal(10,7);not null"`
	Longitude float64   `gorm:"type:decimal(10,7);not null"`
	Label     *string   `gorm:"size:256"`
	Address   *string   `gorm:"size:256"`
	CreatedAt time.Time
}

func (Location) TableName() string {
	return "messages_location"
}
