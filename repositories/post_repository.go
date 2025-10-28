package repositories

import (
	"bifrost/helpers"
	"bifrost/models/post"
	post_payloads "bifrost/models/post/payloads"

	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func (r *PostRepository) DB() *gorm.DB {
	return r.db
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) CreatePost(post *post.Post) error {
	if post.ID == uuid.Nil {
		post.ID = uuid.New()
	}

	// PublicID için Snowflake tarzı ID veya timestamp tabanlı basit artan ID
	if post.PublicID == 0 {
		node, err := helpers.NewNode(1)
		if err != nil {
			return fmt.Errorf("failed to create snowflake node: %w", err)
		}
		post.PublicID = node.Generate().Int64()
	}

	// CreatedAt ve UpdatedAt
	now := time.Now()
	post.CreatedAt = now
	post.UpdatedAt = now

	// GORM ile kaydet
	if err := r.db.Create(post).Error; err != nil {
		return err
	}

	return nil
}

// CreatePoll polls ve seçeneklerini kaydeder
func (r *PostRepository) CreatePoll(poll *post_payloads.Poll) error {
	// Transaction başlat
	fmt.Println("CREATE POLL")
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Poll kaydet
		if err := tx.Create(poll).Error; err != nil {
			return err
		}
		/*
			// PollChoice'ları kaydet
			for i := range poll.Choices {
				poll.Choices[i].PollID = poll.ID
				fmt.Println("ANKET SECIM", poll.Choices[i].Label, poll.Choices[i].ID, poll.ID)
				if err := tx.Create(&poll.Choices[i]).Error; err != nil {
					return err
				}
			}
		*/

		return nil
	})
}

func (r *PostRepository) CreateEvent(event *post_payloads.Event) error {
	return r.db.Create(event).Error
}

func (r *PostRepository) GetPostByID(id uuid.UUID) (*post.Post, error) {
	var p post.Post

	err := r.db.
		Preload("Location").
		Preload("Poll").
		Preload("Poll.Choices").
		Preload("Event").
		Preload("Event.Location").
		Preload("Author").
		Preload("Tags").
		Preload("Attachments").
		First(&p, "id = ?", id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("post with id %s not found", id)
		}
		return nil, err
	}

	return &p, nil
}
