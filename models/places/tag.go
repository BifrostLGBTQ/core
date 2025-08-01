package places

import (
	"github.com/google/uuid"
)

type Tag struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name string    `gorm:"size:64;uniqueIndex;not null"`
	Slug string    `gorm:"size:64;uniqueIndex;not null"`

	Places []Place `gorm:"many2many:place_tags;"`
}

func (Tag) TableName() string {
	return "place_tags"
}
