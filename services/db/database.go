package db

import (
	"bifrost/models"
	"bifrost/models/messages"
	messagePayloads "bifrost/models/messages/payloads"

	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	db.Logger = db.Logger.LogMode(logger.Silent)
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
	db.Exec(`CREATE EXTENSION postgis;`)
	return db.AutoMigrate(
		&models.User{},
		&models.Follow{},
		&models.Like{},
		&models.Block{},
		&models.Favorite{},
		&models.Match{},
		&models.Media{},
		&messages.Message{},
		&messages.Chat{},
		&messages.ChatParticipant{},
		&messages.MessageRead{},
		&messagePayloads.Gift{},
		&messagePayloads.Location{},
		&messagePayloads.File{},
		&messagePayloads.Poll{},
	)
}
