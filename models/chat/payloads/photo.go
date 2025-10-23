package payloads

import (
	"time"

	"bifrost/models/shared"

	"github.com/google/uuid"
)

type Photo struct {
	ID           uuid.UUID           `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	MetadataID   uuid.UUID           `gorm:"type:uuid;not null"`
	FileMetadata shared.FileMetadata `gorm:"foreignKey:MetadataID;references:ID"`

	ThumbnailID   *uuid.UUID           `gorm:"type:uuid"` // opsiyonel thumbnail
	ThumbnailMeta *shared.FileMetadata `gorm:"foreignKey:ThumbnailID;references:ID"`

	CreatedAt time.Time
}

func (Photo) TableName() string {
	return "messages_photo"
}
