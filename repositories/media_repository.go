package repositories

import (
	"bifrost/models/media"
	"bifrost/models/shared"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MediaRepository struct {
	db *gorm.DB
}

func NewMediaRepository(db *gorm.DB) *MediaRepository {
	return &MediaRepository{db: db}
}

func (r *MediaRepository) GenerateStoragePath(ownerID uuid.UUID, ownerType media.OwnerType, role media.MediaRole, filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	id := uuid.New().String()
	date := time.Now().Format("2006-01-02") // YYYY-MM-DD
	baseDir := "./static"

	switch ownerType {
	case media.OwnerUser:
		switch role {
		case media.RoleProfile:
			return fmt.Sprintf("%s/users/%s/avatars/%s%s", baseDir, ownerID.String(), id, ext)
		case media.RoleCover:
			return fmt.Sprintf("%s/users/%s/covers/%s%s", baseDir, ownerID.String(), id, ext)
		default:
			return fmt.Sprintf("%s/users/%s/media/%s%s", baseDir, ownerID.String(), id, ext)
		}
	case media.OwnerPost:
		return fmt.Sprintf("%s/posts/%s/%s/%s%s", baseDir, ownerID.String(), date, id, ext)
	case media.OwnerBlog:
		return fmt.Sprintf("%s/blogs/%s/%s/%s%s", baseDir, ownerID.String(), date, id, ext)
	case media.OwnerChat:
		if role == media.RoleChatVideo {
			return fmt.Sprintf("%s/chat/%s/videos/%s%s", baseDir, ownerID.String(), id, ext)
		}
		return fmt.Sprintf("%s/chat/%s/images/%s%s", baseDir, ownerID.String(), id, ext)
	case media.OwnerPage:
		return fmt.Sprintf("%s/pages/%s/%s/%s%s", baseDir, ownerID.String(), date, id, ext)
	default:
		return fmt.Sprintf("%s/other/%s/%s%s", baseDir, ownerID.String(), id, ext)
	}
}

func (r *MediaRepository) AddUserMedia(db *gorm.DB, userID uuid.UUID, role media.MediaRole, filename, url string, mimeType string, size int64, width, height *int) (*media.Media, error) {
	media := &media.Media{
		ID:        uuid.New(),
		FileID:    uuid.New(), // FileMetadata kaydı için
		OwnerID:   userID,
		OwnerType: media.OwnerUser,
		Role:      role,
		IsPublic:  true,
		File: shared.FileMetadata{
			ID:          uuid.New(),
			URL:         url,
			StoragePath: r.GenerateStoragePath(userID, media.OwnerUser, role, filename),
			MimeType:    mimeType,
			Size:        size,
			Name:        filename,
			Width:       width,
			Height:      height,
			CreatedAt:   time.Now(),
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// DB'ye kaydet
	if err := db.Create(&media.File).Error; err != nil {
		return nil, err
	}

	if err := db.Create(media).Error; err != nil {
		return nil, err
	}

	return media, nil
}

// Generic media ekleme
func (r *MediaRepository) AddMedia(db *gorm.DB, ownerID uuid.UUID, ownerType media.OwnerType, role media.MediaRole, file *multipart.FileHeader) (*media.Media, error) {
	ext := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("%d_%s%s", time.Now().Unix(), uuid.New().String(), ext)
	storagePath := r.GenerateStoragePath(ownerID, ownerType, role, newFileName)

	fmt.Println("STORAGE PATH", storagePath)

	if err := r.SaveUploadedFile(file, storagePath); err != nil {
		return nil, err
	}

	// Burada basit width/height ve duration default null bırakıldı
	media := media.Media{
		ID:        uuid.New(),
		FileID:    uuid.New(),
		OwnerID:   ownerID,
		OwnerType: ownerType,
		Role:      role,
		IsPublic:  true,
		File: shared.FileMetadata{
			ID:          uuid.New(),
			URL:         "/static/" + string(ownerType) + "/" + string(role) + "/" + newFileName,
			StoragePath: storagePath,
			MimeType:    file.Header.Get("Content-Type"),
			Size:        file.Size,
			Name:        file.Filename,
			CreatedAt:   time.Now(),
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.Create(&media.File).Error; err != nil {
		return nil, err
	}
	if err := db.Create(&media).Error; err != nil {
		return nil, err
	}

	return &media, nil
}

// Helper
func (r *MediaRepository) SaveUploadedFile(file *multipart.FileHeader, path string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = dst.ReadFrom(src)
	return err
}

/*


avatarFile := FileMetadata{
	ID:         uuid.New(),
	URL:        "https://cdn.example.com/avatar.png",
	StoragePath: "users/avatars/avatar.png",
	MimeType:   "image/png",
	Size:       120000,
	Width:      ptrInt(512),
	Height:     ptrInt(512),
	CreatedAt:  time.Now(),
}

avatarMedia, _ := mediaRepo.AddMedia(user.ID, OwnerUser, RoleProfile, avatarFile, true)
user.ProfileImageURL = &avatarMedia.File.URL
userRepo.DB().Save(&user)

coverFile := FileMetadata{
	ID:         uuid.New(),
	URL:        "https://cdn.example.com/cover.png",
	StoragePath: "users/covers/cover.png",
	MimeType:   "image/png",
	Size:       240000,
	Width:      ptrInt(1200),
	Height:     ptrInt(400),
	CreatedAt:  time.Now(),
}

coverMedia, _ := mediaRepo.AddMedia(user.ID, OwnerUser, RoleCover, coverFile, true)

chatFile := FileMetadata{
	ID:         uuid.New(),
	URL:        "https://cdn.example.com/chat123.png",
	StoragePath: "chat/room123/2025-10-24/chat123.png",
	MimeType:   "image/png",
	Size:       50000,
	Width:      ptrInt(800),
	Height:     ptrInt(600),
	CreatedAt:  time.Now(),
}

chatMedia, _ := mediaRepo.AddMedia(chatID, OwnerChat, RoleChatImage, chatFile, false)

*/
