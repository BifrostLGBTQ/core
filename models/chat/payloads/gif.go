package payloads

import (
	"time"

	"github.com/google/uuid"
)

type GIF struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	URL       string    `gorm:"size:512;not null"` // GIF dosyasının URL'si
	Width     int       `gorm:"not null"`          // genişlik (pixel)
	Height    int       `gorm:"not null"`          // yükseklik (pixel)
	Duration  int       `gorm:"not null"`          // animasyon süresi (saniye)
	CreatedAt time.Time
}

func (GIF) TableName() string {
	return "messages_gif"
}
