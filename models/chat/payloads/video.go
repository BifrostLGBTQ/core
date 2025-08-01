package payloads

import (
	"time"

	"bifrost/models/shared"

	"github.com/google/uuid"
)

type Video struct {
	ID           uuid.UUID           `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	MetadataID   uuid.UUID           `gorm:"type:uuid;not null"`
	FileMetadata shared.FileMetadata `gorm:"foreignKey:MetadataID;references:ID"`
	Thumbnail    string              `gorm:"size:512"` // küçük önizleme resmi URL'si
	CreatedAt    time.Time
}

func (Video) TableName() string {
	return "messages_videos"
}
