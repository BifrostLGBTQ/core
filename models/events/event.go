package events

import (
	"bifrost/models/places"
	"bifrost/models/shared"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title       string    `gorm:"size:255;not null"`
	Description string    `gorm:"type:text"`
	Location    string    `gorm:"size:255"`
	Latitude    *float64  `gorm:"type:decimal(9,6)"`
	Longitude   *float64  `gorm:"type:decimal(9,6)"`
	StartTime   time.Time `gorm:"not null"`
	EndTime     *time.Time

	Type string `gorm:"size:64;index"`

	HostID   uuid.UUID `gorm:"type:uuid;not null;index"`
	HostType string    `gorm:"size:16;not null"`

	CoverImageID *uuid.UUID
	CoverImage   *shared.FileMetadata `gorm:"foreignKey:CoverImageID"`

	PlaceID *uuid.UUID
	Place   *places.Place `gorm:"foreignKey:PlaceID"`

	Attendees []EventAttendee `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}
