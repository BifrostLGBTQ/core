package post

import (
	"bifrost/models/media"
	"bifrost/models/user"
	"time"

	"bifrost/models/post/payloads"
	"bifrost/models/post/shared"
	global_shared "bifrost/models/shared"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostType string
type ContentCategory string

const (
	PostTypeTimeline   PostType = "timeline"
	PostTypePlace      PostType = "place"
	PostTypeClassified PostType = "classified"
	PostTypeGeneric    PostType = "generic"
	PostTypeNews       PostType = "news" // yeni haber türü
	PostTypeStory      PostType = "story"
)

const (
	ContentNormal       ContentCategory = "normal"       // Standart içerik
	ContentErotic       ContentCategory = "erotic"       // Erotik / yetişkin içerik
	ContentViolence     ContentCategory = "violence"     // Şiddet içerik
	ContentSpam         ContentCategory = "spam"         // Reklam / spam
	ContentPolitical    ContentCategory = "political"    // Politik içerik
	ContentSensitive    ContentCategory = "sensitive"    // Hassas konular (ör: depresyon, travma)
	ContentNSFW         ContentCategory = "nsfw"         // 18+ genel içerik
	ContentSelfPromo    ContentCategory = "self_promo"   // Kendi reklamı / promosyon
	ContentEvent        ContentCategory = "event"        // Etkinlik duyurusu
	ContentAnnouncement ContentCategory = "announcement" // Duyuru
	ContentReview       ContentCategory = "review"       // Yorum / inceleme
	ContentNews         ContentCategory = "news"         // Haber içerik
	ContentArt          ContentCategory = "art"          // Sanat / görsel içerik
	ContentTutorial     ContentCategory = "tutorial"     // Eğitim / rehber
	ContentOther        ContentCategory = "other"        // Diğer
)

type Post struct {
	ID       uuid.UUID  `gorm:"type:uuid;primaryKey"`
	ParentID *uuid.UUID `gorm:"type:uuid;index"`
	Children []Post     `gorm:"foreignKey:ParentID"` // alt postlar
	// Parent post, optional
	PublicID int64 `gorm:"uniqueIndex;not null" json:"public_id"` //snowflake

	AuthorID uuid.UUID `gorm:"type:uuid;index;not null" json:"author_id"`
	Type     PostType  `gorm:"size:50;not null;index"`

	Title   *string                 `gorm:"size:255"`             // optional
	Slug    *string                 `gorm:"size:255;uniqueIndex"` // optional
	Content *shared.LocalizedString `gorm:"type:jsonb"`           // optional {"en":"Hello","tr":"Merhaba"}
	Summary *shared.LocalizedString `gorm:"type:jsonb"`           // optional {"en":"Short desc","tr":"Kısa açıklama"}

	Published   bool       `gorm:"default:false;index"`
	PublishedAt *time.Time `gorm:"index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt  `gorm:"index"`
	Extras      *map[string]any `gorm:"type:jsonb"` // Universal flexible extra fields
	Author      user.User       `gorm:"foreignKey:AuthorID;references:ID"`
	Tags        []payloads.Tag  `gorm:"many2many:post_tags;"`
	Attachments []*media.Media  `gorm:"polymorphic:Owner;polymorphicValue:post;constraint:OnDelete:CASCADE" json:"media,omitempty"`

	Poll  *payloads.Poll  `gorm:"polymorphic:Contentable;constraint:OnDelete:CASCADE"`
	Event *payloads.Event `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"` // One-to-one event

	Location *global_shared.Location `gorm:"polymorphic:Contentable;constraint:OnDelete:CASCADE"`
}

func (Post) TableName() string {
	return "posts"
}
