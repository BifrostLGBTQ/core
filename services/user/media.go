package services

import (
	"bifrost/models/media"
	"bifrost/repositories"
	"mime/multipart"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MediaService struct {
	repo *repositories.MediaRepository

	db *gorm.DB
}

func NewMediaService(db *gorm.DB) *MediaService {
	return &MediaService{db: db}
}

func (s *MediaService) AddMedia(ownerID uuid.UUID, ownerType media.OwnerType, role media.MediaRole, file *multipart.FileHeader) (*media.Media, error) {
	return s.repo.AddMedia(s.db, ownerID, ownerType, role, file)
}
