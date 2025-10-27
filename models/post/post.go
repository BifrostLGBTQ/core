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
	ID       uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	ParentID *uuid.UUID `gorm:"type:uuid;index" json:"parent_id,omitempty"`
	Children []Post     `gorm:"foreignKey:ParentID" json:"children,omitempty"`

	PublicID int64 `gorm:"uniqueIndex;not null" json:"public_id"`

	AuthorID uuid.UUID `gorm:"type:uuid;index;not null" json:"author_id"`
	Type     PostType  `gorm:"size:50;not null;index" json:"type"`

	Title   *string                 `gorm:"size:255" json:"title,omitempty"`
	Slug    *string                 `gorm:"size:255;uniqueIndex" json:"slug,omitempty"`
	Content *shared.LocalizedString `gorm:"type:jsonb" json:"content,omitempty"`
	Summary *shared.LocalizedString `gorm:"type:jsonb" json:"summary,omitempty"`

	Published   bool           `gorm:"default:false;index" json:"published"`
	PublishedAt *time.Time     `gorm:"index" json:"published_at,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Extras *map[string]any `gorm:"type:jsonb" json:"extras,omitempty"`

	Author      user.User      `gorm:"foreignKey:AuthorID;references:ID" json:"author"`
	Tags        []payloads.Tag `gorm:"many2many:post_tags;" json:"tags,omitempty"`
	Attachments []*media.Media `gorm:"polymorphic:Owner;polymorphicValue:post;constraint:OnDelete:CASCADE" json:"attachments,omitempty"`

	Poll     *payloads.Poll          `gorm:"polymorphic:Contentable;constraint:OnDelete:CASCADE" json:"poll,omitempty"`
	Event    *payloads.Event         `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE" json:"event,omitempty"`
	Location *global_shared.Location `gorm:"polymorphic:Contentable;constraint:OnDelete:CASCADE" json:"location,omitempty"`
}

func (Post) TableName() string {
	return "posts"
}
