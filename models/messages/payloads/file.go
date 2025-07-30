package payloads

import "time"

type File struct {
	ID        uint   `gorm:"primaryKey"`
	URL       string `gorm:"size:512;not null"`
	FileType  string `gorm:"size:64"` // mime tipi veya türü
	Size      int64
	Name      string `gorm:"size:256"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
