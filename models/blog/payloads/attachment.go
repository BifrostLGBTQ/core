package payloads

import (
	"time"

	"bifrost/models/shared"

	"github.com/google/uuid"
)

type Attachment struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	BlogPostID uuid.UUID `gorm:"type:uuid;index;not null"`
	FileID     uuid.UUID `gorm:"type:uuid;not null;index"`
	Type       string    `gorm:"size:32;not null"` // "image", "video", "audio", "document" vs.
	Caption    *string   `gorm:"size:255"`         // Açıklama metni
	OrderIndex int       `gorm:"default:0"`        // Sıralama
	CreatedAt  time.Time

	FileMetadata shared.FileMetadata `gorm:"foreignKey:FileID;references:ID;constraint:OnDelete:CASCADE"`
}

func (Attachment) TableName() string {
	return "blog_post_attachments"
}
