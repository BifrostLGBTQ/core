package chat

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"bifrost/models"
	message_payloads "bifrost/models/chat/payloads"
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

	Gift     *message_payloads.Gift     `gorm:"foreignKey:PayloadID;references:ID"`
	Location *message_payloads.Location `gorm:"foreignKey:PayloadID;references:ID"`
	File     *message_payloads.File     `gorm:"foreignKey:PayloadID;references:ID"`
	Poll     *message_payloads.Poll     `gorm:"foreignKey:PayloadID;references:ID"`
	GIF      *message_payloads.GIF      `gorm:"foreignKey:PayloadID;references:ID"`
	Photo    *message_payloads.Photo    `gorm:"foreignKey:PayloadID;references:ID"`
	Video    *message_payloads.Video    `gorm:"foreignKey:PayloadID;references:ID"`
	Audio    *message_payloads.Audio    `gorm:"foreignKey:PayloadID;references:ID"`
	Sticker  *message_payloads.Sticker  `gorm:"foreignKey:PayloadID;references:ID"`
	Call     *message_payloads.Call     `gorm:"foreignKey:PayloadID;references:ID"`
	System   *message_payloads.System   `gorm:"foreignKey:PayloadID;references:ID"`

	Reads []MessageRead `gorm:"foreignKey:MessageID"`

	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Message) TableName() string {
	return "messages"
}
