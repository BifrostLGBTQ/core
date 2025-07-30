package payloads

import (
	"time"

	"github.com/google/uuid"
)

type Video struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	URL       string    `gorm:"size:512;not null"`
	Duration  int       `gorm:"not null"` // saniye cinsinden
	Thumbnail string    `gorm:"size:512"` // küçük önizleme resmi URL'si
	CreatedAt time.Time
}

func (Video) TableName() string {
	return "messages_videos"
}
