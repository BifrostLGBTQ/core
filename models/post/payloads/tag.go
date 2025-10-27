package payloads

import (
	"github.com/google/uuid"
)

type Tag struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name     string    `gorm:"size:64;uniqueIndex;not null"`
	UseCount int64     `gorm:"default:0"` // kaç kez kullanıldığı
}

func (Tag) TableName() string {
	return "tags"
}
