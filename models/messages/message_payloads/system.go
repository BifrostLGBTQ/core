package message_payloads

import (
	"time"

	"github.com/google/uuid"
)

type SystemMessageType int

const (
	SYSMSG_UNKNOWN SystemMessageType = iota
	SYSMSG_USER_JOINED
	SYSMSG_USER_LEFT
	SYSMSG_USER_REMOVED
	SYSMSG_GROUP_RENAMED
	SYSMSG_MESSAGE_PINNED
	SYSMSG_MESSAGE_UNPINNED
	SYSMSG_CALL_STARTED
	SYSMSG_CALL_ENDED
)

// Sistem mesajları için payload
type System struct {
	ID          uuid.UUID         `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Type        SystemMessageType `gorm:"not null"`
	Description string            `gorm:"size:512"`  // İsteğe bağlı açıklama veya içerik
	TargetUser  *uuid.UUID        `gorm:"type:uuid"` // örneğin USER_JOINED için ilgili kullanıcı
	Metadata    string            `gorm:"type:text"` // JSON olarak ek veri (isteğe bağlı)
	CreatedAt   time.Time
}

func (System) TableName() string {
	return "messages_system"
}
