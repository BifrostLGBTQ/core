package payloads

import (
	"github.com/google/uuid"
)

type SexualOrientation struct {
	ID           uuid.UUID                       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Key          string                          `gorm:"uniqueIndex;not null" json:"key"`
	Order        int                             `json:"order"`
	Translations []*SexualOrientationTranslation `gorm:"foreignKey:OrientationID" json:"translations,omitempty"`
}

type SexualOrientationTranslation struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	OrientationID uuid.UUID `gorm:"type:uuid;index;not null" json:"orientation_id"`
	Language      string    `gorm:"type:varchar(5);not null" json:"language"`
	Label         string    `gorm:"type:text;not null" json:"label"`
}
