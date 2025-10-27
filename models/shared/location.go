// models/payloads/location.go
package shared

import (
	"bifrost/extensions"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	LocationOwnerPost  = "post"
	LocationOwnerEvent = "event"
	LocationOwnerUser  = "user"
)

// OwnerType: örn "post", "event", "user", "chat", ...
type Location struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ContentableID   uuid.UUID
	ContentableType string
	Address         *string `gorm:"size:512"`
	City            *string `gorm:"size:128"`
	Country         *string `gorm:"size:128"`
	Postal          *string `gorm:"size:32"`
	// latitude/longitude fallback (if you don't use PostGIS)
	Latitude  *float64 `gorm:"type:numeric(10,6)"`
	Longitude *float64 `gorm:"type:numeric(10,6)"`
	// If PostGIS available, store geography point for fast distance queries
	// Use PostGIS column type via tag (GORM may need custom datatype registration)
	LocationPoint *extensions.PostGISPoint `gorm:"type:geography(Point,4326)" json:"location_point"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Location) TableName() string {
	return "locations"
}
