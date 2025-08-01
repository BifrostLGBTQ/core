package blog

import (
	"time"

	"bifrost/models"
	"bifrost/models/blog/payloads"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlogPost struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey"`
	AuthorID    uuid.UUID  `gorm:"type:uuid;index;not null"`
	Title       string     `gorm:"size:255;not null"`
	Slug        string     `gorm:"size:255;uniqueIndex;not null"`
	Content     string     `gorm:"type:text;not null"`
	Summary     string     `gorm:"type:text"`
	CoverImage  *string    `gorm:"size:512"`
	Published   bool       `gorm:"default:false;index"`
	PublishedAt *time.Time `gorm:"index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	Author      models.User           `gorm:"foreignKey:AuthorID;references:ID"`
	Tags        []payloads.Tag        `gorm:"many2many:blogpost_tags;"`
	Attachments []payloads.Attachment `gorm:"foreignKey:BlogPostID;constraint:OnDelete:CASCADE"`
}

func (BlogPost) TableName() string {
	return "blog_posts"
}
