package message_payloads

import (
	"time"

	"bifrost/models/shared"

	"github.com/google/uuid"
)

type Audio struct {
	ID           uuid.UUID           `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	MetadataID   uuid.UUID           `gorm:"type:uuid;not null"`
	FileMetadata shared.FileMetadata `gorm:"foreignKey:MetadataID;references:ID"`
	CreatedAt    time.Time
}

func (Audio) TableName() string {
	return "messages_audio"
}
