package chat

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Chat struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Type        ChatType  `gorm:"index;not null"` // private, group, channel
	Title       *string   `gorm:"size:128"`       // grup/kanal adı
	Description *string   `gorm:"size:512"`
	AvatarURL   *string
	CreatorID   uuid.UUID  `gorm:"type:uuid;index;not null"` // UUID olmalı
	PinnedMsgID *uuid.UUID `gorm:"type:uuid;index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	Participants []ChatParticipant
	Messages     []Message
}

func (Chat) TableName() string {
	return "chats"
}
