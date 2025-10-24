package payloads

import (
	"github.com/google/uuid"
)

type Fantasy struct {
	ID           uuid.UUID             `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Category     string                `gorm:"type:varchar(50);not null" json:"category"`
	Translations []*FantasyTranslation `gorm:"foreignKey:FantasyID" json:"translations,omitempty"`
}

type FantasyTranslation struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	FantasyID   uuid.UUID `gorm:"type:uuid;index;not null" json:"fantasy_id"`
	Language    string    `gorm:"type:varchar(10);not null" json:"language"`
	Label       string    `gorm:"type:text;not null" json:"label"`
	Description string    `gorm:"type:text" json:"description,omitempty"`
}

type UserFantasy struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;index;not null" json:"user_id"`
	FantasyID uuid.UUID `gorm:"type:uuid;index;not null" json:"fantasy_id"`
	Notes     *string   `gorm:"type:text" json:"notes,omitempty"`

	Fantasy *Fantasy `gorm:"foreignKey:FantasyID;references:ID" json:"fantasy,omitempty"`
}

func (UserFantasy) TableName() string {
	return "user_fantasies"
}

func (FantasyTranslation) TableName() string {
	return "fantasy_translations"
}

func (Fantasy) TableName() string {
	return "fantasies"
}
