package seeders

import (
	"encoding/json"
	"fmt"
	"os"

	"bifrost/models/user/payloads"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedFantasies(db *gorm.DB) error {
	// JSON dosyasını aç
	file, err := os.Open("static/data/sexual_preferences.json")
	if err != nil {
		return fmt.Errorf("Cannot open JSON file: %w", err)
	}
	defer file.Close()

	var data []struct {
		Label       string `json:"label"`
		Description string `json:"description"`
		Category    string `json:"category"`
	}
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return fmt.Errorf("Cannot decode JSON: %w", err)
	}

	for _, item := range data {
		var fantasy payloads.Fantasy
		// Eğer kategori ve label mevcutsa atla
		err := db.Joins("JOIN fantasy_translations ON fantasy_translations.fantasy_id = fantasies.id").
			Where("fantasy_translations.label = ? AND fantasies.category = ?", item.Label, item.Category).
			First(&fantasy).Error
		if err == gorm.ErrRecordNotFound {
			// Yeni fantasy ekle
			fantasy = payloads.Fantasy{
				ID:       uuid.New(),
				Category: item.Category,
				Translations: []*payloads.FantasyTranslation{
					{
						ID:          uuid.New(),
						Language:    "en",
						Label:       item.Label,
						Description: item.Description,
					},
				},
			}
			if err := db.Create(&fantasy).Error; err != nil {
				fmt.Printf("Failed to insert fantasy '%s': %v\n", item.Label, err)
			}
		}
	}
	return nil
}
