package messages

import (
	"time"

	"bifrost/models"
	"bifrost/models/messages/payloads"

	"gorm.io/gorm"
)

type MessageType string
type ParticipantRole string
type MessageStatus string

const (
	Text     MessageType = "text"
	Image    MessageType = "image"
	Video    MessageType = "video"
	Audio    MessageType = "audio"
	GIF      MessageType = "gif"
	Sticker  MessageType = "sticker"
	File     MessageType = "file"
	Location MessageType = "location"
	System   MessageType = "system"
	Gift     MessageType = "gift" // hediye tipi
	Poll     MessageType = "poll" // anket
)

const (
	Pending   MessageStatus = "pending"
	Delivered MessageStatus = "delivered"
	Seen      MessageStatus = "seen"
	Deleted   MessageStatus = "deleted"
)

type ChatType string

const (
	Private ChatType = "private"
	Group   ChatType = "group"
	Channel ChatType = "channel"
)

type MessageRead struct {
	ID        uint `gorm:"primaryKey"`
	MessageID uint `gorm:"index;not null"`
	Message   Message
	UserID    uint `gorm:"index;not null"`
	User      models.User
	ReadAt    time.Time `gorm:"autoCreateTime"`
}

type Chat struct {
	ID          uint     `gorm:"primaryKey"`
	Type        ChatType `gorm:"index;not null"` // private, group, channel
	Title       *string  `gorm:"size:128"`       // grup/kanal adı
	Description *string  `gorm:"size:512"`
	AvatarURL   *string
	CreatorID   uint `gorm:"index;not null"`
	PinnedMsgID *uint
	PinnedMsg   *Message `gorm:"foreignKey:PinnedMsgID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	Participants []ChatParticipant
	Messages     []Message
}

type ChatParticipant struct {
	ID        uint `gorm:"primaryKey"`
	ChatID    uint `gorm:"index;not null"`
	Chat      Chat
	UserID    uint `gorm:"index;not null"`
	User      models.User
	Role      ParticipantRole `gorm:"size:32;default:'member'"`
	IsMuted   bool            `gorm:"default:false"`
	JoinedAt  time.Time
	LeftAt    *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Message struct {
	ID     uint `gorm:"primaryKey"`
	ChatID uint `gorm:"index;not null"`
	Chat   Chat

	SenderID uint `gorm:"index;not null"`
	Sender   models.User

	ReplyToID *uint    `gorm:"index"`
	ReplyTo   *Message `gorm:"foreignKey:ReplyToID"`

	ForwardedFromID *uint
	ForwardedFrom   *models.User

	Type MessageType `gorm:"size:16;index;not null"` // text, image, gif, gift vs.

	Content string `gorm:"type:text"` // metin içeriği (caption, text, vb.)

	// Polimorfik Payload: her mesaj sadece bir tip detay içerir
	PayloadType *string `gorm:"size:32;index"` // "gift", "location", "file", "poll"
	PayloadID   *uint   `gorm:"index"`

	Status MessageStatus `gorm:"size:16;default:'delivered';index"`

	IsSystem bool `gorm:"default:false"`
	IsPinned bool `gorm:"default:false"`

	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Preload ilişkileri:
	Gift     *payloads.Gift     `gorm:"foreignKey:PayloadID;references:ID"`
	Location *payloads.Location `gorm:"foreignKey:PayloadID;references:ID"`
	File     *payloads.File     `gorm:"foreignKey:PayloadID;references:ID"`
	Poll     *payloads.Poll     `gorm:"foreignKey:PayloadID;references:ID"`

	Reads []MessageRead `gorm:"foreignKey:MessageID"`
}
