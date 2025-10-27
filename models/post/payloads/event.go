package payloads

import (
	"time"

	"bifrost/models/post/shared"
	global_shared "bifrost/models/shared"

	"github.com/google/uuid"
)

type EventAttendee struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	EventID uuid.UUID `gorm:"type:uuid;not null;index"`
	UserID  uuid.UUID `gorm:"type:uuid;not null;index"`

	Status   string    `gorm:"size:32;default:'interested'"` // "going", "interested", "invited", "declined"
	JoinedAt time.Time `gorm:"autoCreateTime"`
}

type Event struct {
	ID          uuid.UUID              `gorm:"type:uuid;primaryKey"`
	PostID      uuid.UUID              `gorm:"type:uuid;uniqueIndex;not null"` // Post ile birebir ilişki
	Title       shared.LocalizedString `gorm:"type:jsonb"`
	Description shared.LocalizedString `gorm:"type:jsonb"`
	StartTime   *time.Time
	EndTime     *time.Time

	// location
	LocationID *uuid.UUID `gorm:"type:uuid"`
	//Location   *global_shared.Location `gorm:"foreignKey:OwnerID;references:ID;->;constraint:OnDelete:SET NULL"`
	Location *global_shared.Location `gorm:"polymorphic:Contentable;constraint:OnDelete:CASCADE"`

	Type string `gorm:"size:64;index"`

	Attendees []EventAttendee `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func (Event) TableName() string {
	return "events"
}

func (EventAttendee) TableName() string {
	return "event_attendees"
}
