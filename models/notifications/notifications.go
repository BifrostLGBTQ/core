package notifications

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID uuid.UUID `gorm:"type:uuid;index;not null"` // Bildirimin hedef kullanıcısı

	// Tip: "chat_message", "friend_request", "event_reminder", "system_alert", vb.
	Type string `gorm:"size:64;index;not null"`

	// Bildirim başlığı ve kısa metni (önizleme)
	Title   string `gorm:"size:255"`
	Message string `gorm:"type:text"`

	// Ek detaylar, örn. chat mesajının içeriği, link, id vb.
	Payload map[string]interface{} `gorm:"type:jsonb"` // PostgreSQL JSONB, MySQL için json

	IsRead    bool      `gorm:"default:false;index"` // Okundu mu?
	IsShown   bool      `gorm:"default:false"`       // Kullanıcıya gösterildi mi? (push notification gibi)
	CreatedAt time.Time `gorm:"autoCreateTime;index"`
	ReadAt    *time.Time
	ShownAt   *time.Time
	DeletedAt *time.Time `gorm:"index"`
}
