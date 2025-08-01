package shared

import (
	"time"

	"github.com/google/uuid"
)

type FileMetadata struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	URL       string    `gorm:"size:512;not null"` // CDN veya public URL
	MimeType  string    `gorm:"size:128;not null"` // "image/png", "video/mp4" vs.
	Size      int64     `gorm:"not null"`          // Bytes cinsinden
	Name      string    `gorm:"size:255"`          // Orijinal dosya adı
	Width     *int      `gorm:"null"`              // Resim/video için
	Height    *int      `gorm:"null"`              // Resim/video için
	Duration  *float64  `gorm:"null"`              // Ses/video için saniye cinsinden
	CreatedAt time.Time
}

func (FileMetadata) TableName() string {
	return "file_metadata"
}
