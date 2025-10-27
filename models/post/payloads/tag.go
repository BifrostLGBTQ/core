package payloads

import (
	"github.com/google/uuid"
)

type Tag struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name     string    `gorm:"size:64;uniqueIndex;not null" json:"name"`
	UseCount int64     `gorm:"default:0" json:"use_count"` // kaç kez kullanıldığı
}

func (Tag) TableName() string {
	return "tags"
}
