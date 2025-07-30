package messages

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"bifrost/models"
	"bifrost/models/messages/payloads"
)

type Message struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	ChatID   uuid.UUID `gorm:"type:uuid;index;not null"`
	SenderID uuid.UUID `gorm:"type:uuid;index;not null"`

	ReplyToID       *uuid.UUID
	ForwardedFromID *uuid.UUID

	Type        MessageType   `gorm:"type:varchar(16);index;not null"`
	Content     string        `gorm:"type:text"`
	PayloadType *string       `gorm:"size:32;index"`
	PayloadID   *uuid.UUID    `gorm:"type:uuid;index"`
	Status      MessageStatus `gorm:"type:varchar(16);default:'delivered';index"`
	IsSystem    bool
	IsPinned    bool

	// Relations
	Chat          Chat
	Sender        models.User
	ReplyTo       *Message `gorm:"foreignKey:ReplyToID"`
	ForwardedFrom *models.User

	Gift     *payloads.Gift     `gorm:"foreignKey:PayloadID;references:ID"`
	Location *payloads.Location `gorm:"foreignKey:PayloadID;references:ID"`
	File     *payloads.File     `gorm:"foreignKey:PayloadID;references:ID"`
	Poll     *payloads.Poll     `gorm:"foreignKey:PayloadID;references:ID"`
	GIF      *payloads.GIF      `gorm:"foreignKey:PayloadID;references:ID"`
	Photo    *payloads.Photo    `gorm:"foreignKey:PayloadID;references:ID"`
	Video    *payloads.Video    `gorm:"foreignKey:PayloadID;references:ID"`
	Audio    *payloads.Audio    `gorm:"foreignKey:PayloadID;references:ID"`

	Reads []MessageRead `gorm:"foreignKey:MessageID"`

	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
