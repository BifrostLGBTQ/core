package payloads

import "time"

type Gift struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:64"`
	Emoji     string `gorm:"size:8"`
	MediaURL  string `gorm:"size:256"` // Gif, video, sticker url olabilir
	CreatedAt time.Time
	UpdatedAt time.Time
}
