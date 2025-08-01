package chat

import (
	"time"

	"github.com/google/uuid"

	"bifrost/models"
)

type ChatParticipant struct {
	ID       uuid.UUID       `gorm:"type:uuid;primaryKey"`
	ChatID   uuid.UUID       `gorm:"type:uuid;index;not null"`
	UserID   uuid.UUID       `gorm:"type:uuid;index;not null"`
	Role     ParticipantRole `gorm:"type:varchar(32);default:'member'"`
	IsMuted  bool
	JoinedAt time.Time
	LeftAt   *time.Time

	Chat Chat
	User models.User

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ChatParticipant) TableName() string {
	return "messages_chat_participants"
}
