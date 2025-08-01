package places

import (
	"time"

	"bifrost/models/shared"

	"github.com/google/uuid"
)

type PlaceVisit struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index"` // Ziyaret eden kullanıcı
	PlaceID   uuid.UUID `gorm:"type:uuid;not null;index"` // Ziyaret edilen mekan
	VisitedAt time.Time `gorm:"not null;index"`           // Ziyaret zamanı

	Comment *string `gorm:"type:text"` // Opsiyonel kullanıcı yorumu
	Rating  *int    `gorm:"type:int"`  // 1-5 arası puan, opsiyonel

	// Fotoğraf yerine shared.FileMetadata kullanımı
	PhotoID *uuid.UUID           `gorm:"type:uuid;index"` // Opsiyonel dosya
	Photo   *shared.FileMetadata `gorm:"foreignKey:PhotoID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
