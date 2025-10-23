package payloads

import (
	"github.com/google/uuid"
)

type Fantasy struct {
	ID           uuid.UUID             `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Category     string                `gorm:"type:varchar(50);not null"`
	Translations []*FantasyTranslation `gorm:"foreignKey:FantasyID"`
}

type FantasyTranslation struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FantasyID   uuid.UUID `gorm:"type:uuid;index;not null"`
	Language    string    `gorm:"type:varchar(10);not null"`
	Label       string    `gorm:"type:text;not null"`
	Description string    `gorm:"type:text"`
}

type UserFantasy struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;index;not null"`
	FantasyID uuid.UUID `gorm:"type:uuid;index;not null"`
	Notes     *string   `gorm:"type:text"`

	Fantasy *Fantasy `gorm:"foreignKey:FantasyID"`
}
