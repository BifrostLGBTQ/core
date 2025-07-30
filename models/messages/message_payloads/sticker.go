package message_payloads

import (
	"time"

	"github.com/google/uuid"
)

type Sticker struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FileID  uuid.UUID `gorm:"type:uuid;not null;index"` // sticker dosyasına ait media ID
	SetName *string   `gorm:"size:128"`                 // sticker set ismi (varsa)
	Emoji   *string   `gorm:"size:16"`                  // sticker'a eşlik eden emoji (örn: 😄)

	Width    *int
	Height   *int
	Format   *string `gorm:"size:32"` // webp, png, svg vs.
	MimeType *string `gorm:"size:64"`

	CreatedAt time.Time
}

func (Sticker) TableName() string {
	return "messages_stickers"
}
