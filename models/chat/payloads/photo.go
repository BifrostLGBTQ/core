package payloads

import (
	"time"

	"bifrost/models/shared"

	"github.com/google/uuid"
)

type Photo struct {
	ID           uuid.UUID           `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FileMetadata shared.FileMetadata `gorm:"foreignKey:MetadataID;references:ID"`
	MetadataID   uuid.UUID           `gorm:"type:uuid;not null"`
	CreatedAt    time.Time
}

func (Photo) TableName() string {
	return "messages_photo"
}
