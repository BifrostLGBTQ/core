package payloads

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	URL         string    `gorm:"not null"`
	FileName    string    `gorm:"size:256"`
	ContentType string    `gorm:"size:128"`
	Size        int64     `gorm:"not null"` // bytes
	CreatedAt   time.Time
}
