package payloads

import (
	"github.com/google/uuid"
)

type Tag struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name string    `gorm:"size:64;uniqueIndex;not null"`
}

func (Tag) TableName() string {
	return "tags"
}
