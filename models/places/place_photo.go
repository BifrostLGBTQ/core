package places

import (
	"time"

	shared "bifrost/models/shared"

	"github.com/google/uuid"
)

type PlacePhoto struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	PlaceID   uuid.UUID `gorm:"type:uuid;index;not null"`
	FileID    uuid.UUID `gorm:"type:uuid;not null"`
	IsCover   bool      `gorm:"default:false"` // Kapak fotoğrafı mı?
	SortOrder int       `gorm:"default:0"`     // Sıralama için

	Place Place               `gorm:"foreignKey:PlaceID;references:ID"`
	File  shared.FileMetadata `gorm:"foreignKey:FileID;references:ID"`

	UploadedAt time.Time
}

func (PlacePhoto) TableName() string {
	return "place_photos"
}
