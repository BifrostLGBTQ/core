package shared

import (
	"time"

	"github.com/google/uuid"
)

type FileMetadata struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	URL         string    `gorm:"size:512;not null" json:"url"`          // CDN veya public URL
	StoragePath string    `gorm:"size:512;not null" json:"storage_path"` // CDN veya local path
	MimeType    string    `gorm:"size:128;not null" json:"mime_type"`    // "image/png", "video/mp4" vs.
	Size        int64     `gorm:"not null" json:"size"`                  // Bytes cinsinden
	Name        string    `gorm:"size:255" json:"name"`                  // Orijinal dosya adı
	Width       *int      `gorm:"null" json:"width,omitempty"`           // Resim/video için
	Height      *int      `gorm:"null" json:"height,omitempty"`          // Resim/video için
	Duration    *float64  `gorm:"null" json:"duration,omitempty"`        // Ses/video için saniye cinsinden
	CreatedAt   time.Time `json:"created_at"`                            // Oluşturulma zamanı
}

func (FileMetadata) TableName() string {
	return "file_metadata"
}
