package places

import (
	shared "bifrost/models/shared"
	"time"

	"github.com/google/uuid"
)

type Place struct {
	ID             uuid.UUID        `gorm:"type:uuid;primaryKey"`
	Name           string           `gorm:"size:255;not null"`
	Description    string           `gorm:"type:text"`
	PlaceType      string           `gorm:"size:64;index;not null"` // "hotel", "restaurant", "cafe", "bar", "store" vs.
	Address        string           `gorm:"size:512"`
	Latitude       float64          `gorm:"not null"`
	Longitude      float64          `gorm:"not null"`
	Website        *string          `gorm:"size:255"`
	Phone          *string          `gorm:"size:64"`
	IsLgbtFriendly bool             `gorm:"default:false;index"`
	CreatorID      *uuid.UUID       `gorm:"type:uuid;index"` // Topluluk eklerse, kim ekledi
	Tags           []Tag            `gorm:"many2many:place_tags;"`
	Photos         []PlacePhoto     `gorm:"foreignKey:PlaceID"`
	Visits         []PlaceVisit     `gorm:"foreignKey:PlaceID"`
	Comments       []shared.Comment `gorm:"foreignKey:TargetID;constraint:OnDelete:CASCADE;"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
