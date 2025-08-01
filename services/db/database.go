package db

import (
	"bifrost/models"
	"bifrost/models/chat"
	message_payloads "bifrost/models/chat/payloads"

	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Global değişken olarak veritabanı bağlantısı

func InitDB() error {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		panic("DATABASE_URL is required")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		// Hata işleme
	}

	sqlDB.SetMaxIdleConns(10)           // Boşta bekleyen bağlantıların maksimum sayısı
	sqlDB.SetMaxOpenConns(0)            // Aynı anda açık olabilecek maksimum bağlantı sayısı
	sqlDB.SetConnMaxLifetime(time.Hour) // Bağlantının yeniden kullanılabilir olacağı maksimum süre

	DB = db
	return nil
}

func Migrate(db *gorm.DB) error {
	fmt.Println("Migration:Begin")
	//db.Logger = db.Logger.LogMode(logger.Silent)
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
	db.Exec(`CREATE EXTENSION postgis;`)

	err := db.AutoMigrate(
		&models.User{},
		&models.Follow{},
		&models.Like{},
		&models.Block{},
		&models.Favorite{},
		&models.Match{},
		&models.Media{},

		// öncelikle payload'lar
		&message_payloads.Gift{},
		&message_payloads.Location{},
		&message_payloads.File{},
		&message_payloads.Poll{},
		&message_payloads.PollOption{},
		&message_payloads.PollVote{},
		&message_payloads.GIF{},
		&message_payloads.Photo{},
		&message_payloads.Video{},
		&message_payloads.Audio{},
		&message_payloads.Sticker{},
		&message_payloads.Call{},
		&message_payloads.System{},
		// sonra mesajlar
		&chat.Message{},
		&chat.Chat{},
		&chat.ChatParticipant{},
		&chat.MessageRead{},
	)

	db.Exec(`
	DO $$
	BEGIN
		IF NOT EXISTS (
			SELECT 1 
			FROM pg_constraint 
			WHERE conname = 'fk_chats_pinned_msg'
		) THEN
			ALTER TABLE chats 
			ADD CONSTRAINT fk_chats_pinned_msg 
			FOREIGN KEY (pinned_msg_id) REFERENCES messages(id);
		END IF;
	END
	$$;
	`)

	return err
}
