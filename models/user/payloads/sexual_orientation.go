package payloads

import (
	"github.com/google/uuid"
)

type SexualOrientation struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Key          string    `gorm:"uniqueIndex;not null"`
	Order        int
	Translations []SexualOrientationTranslation `gorm:"foreignKey:OrientationID"`
}

type SexualOrientationTranslation struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OrientationID uuid.UUID `gorm:"type:uuid;index;not null"`
	Language      string    `gorm:"type:varchar(5);not null"`
	Label         string    `gorm:"type:text;not null"`
}
