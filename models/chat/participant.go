package chat

import (
	"bifrost/models/user"
	"time"

	"github.com/google/uuid"
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
	User user.User

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ChatParticipant) TableName() string {
	return "messages_chat_participants"
}
