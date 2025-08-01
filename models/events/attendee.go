package events

import (
	"time"

	"github.com/google/uuid"
)

type EventAttendee struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	EventID uuid.UUID `gorm:"type:uuid;not null;index"`
	UserID  uuid.UUID `gorm:"type:uuid;not null;index"`

	Status   string    `gorm:"size:32;default:'interested'"` // "going", "interested", "invited", "declined"
	JoinedAt time.Time `gorm:"autoCreateTime"`

	// Composite unique constraint
	// (bir kullanıcı bir etkinliğe sadece 1 kez katılabilir)
}

func (EventAttendee) TableName() string {
	return "event_attendees"
}
