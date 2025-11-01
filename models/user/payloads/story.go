package payloads

import (
	"bifrost/models/media"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Story struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`

	MediaID uuid.UUID    `gorm:"type:uuid;not null" json:"media_id"`
	Media   *media.Media `gorm:"constraint:OnDelete:CASCADE;foreignKey:MediaID;references:ID" json:"media"`

	Caption    *string   `gorm:"type:text" json:"caption,omitempty"`
	ExpiresAt  time.Time `gorm:"index" json:"expires_at"`
	IsExpired  bool      `gorm:"default:false" json:"is_expired"`
	IsArchived bool      `gorm:"default:false" json:"is_archived"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
