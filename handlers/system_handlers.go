package handlers

import (
	"bifrost/models/user/payloads"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

// TranslationMap map[string]Translation
type TranslationMap map[string]FantasyTranslationResponse

type FantasyTranslationResponse struct {
	Label       string `json:"label"`
	Description string `json:"description"`
}

// FantasyResponse tek fantezi objesi
type FantasyResponse struct {
	ID           string         `json:"id"`
	Category     string         `json:"category"`
	Translations TranslationMap `json:"translations"`
}

// CountryResponse tek ülke objesi
type CountryResponse struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type LanguageResponse struct {
	Code string `json:"code"`
	Flag string `json:"flag"`
	Name string `json:"name"`
}

// InitialData dönecek ana struct
type InitialData struct {
	Fantasies          map[string]FantasyResponse   `json:"fantasies"`
	Countries          map[string]CountryResponse   `json:"countries"`
	SexualOrientations map[string]map[string]string `json:"sexual_orientations"` // key -> {lang -> label}
	Languages          map[string]LanguageResponse  `json:"languages"`

	Status string `json:"status"`
}

func HandleInitialSync(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Tüm fantezileri çek
		var fantasies []payloads.Fantasy
		if err := db.Preload("Translations").Find(&fantasies).Error; err != nil {
			http.Error(w, "Failed to fetch fantasies", http.StatusInternalServerError)
			return
		}

		// 2. Fantezileri map formatına çevir
		fantasyMap := make(map[string]FantasyResponse)
		for _, f := range fantasies {
			translations := make(TranslationMap)
			for _, t := range f.Translations {
				lang := t.Language
				if lang == "" {
					lang = "en" // default İngilizce
				}
				translations[lang] = FantasyTranslationResponse{
					Label:       t.Label,
					Description: t.Description,
				}
			}

			fantasyMap[f.ID.String()] = FantasyResponse{
				ID:           f.ID.String(),
				Category:     f.Category,
				Translations: translations,
			}
		}

		// 3. Ülkeleri çek
		// Örneğin countries tablosu veya sabit listeden
		countries := map[string]CountryResponse{
			"TR": {Code: "TR", Name: "Turkey"},
			"US": {Code: "US", Name: "United States"},
			// dilediğin kadar ekle
		}

		// Languages
		languages := map[string]LanguageResponse{
			"en": {Code: "en", Flag: "🇺🇸", Name: "English"},
			"tr": {Code: "tr", Flag: "🇹🇷", Name: "Türkçe"},
			"es": {Code: "es", Flag: "🇪🇸", Name: "Español"},
			"he": {Code: "he", Flag: "🇮🇱", Name: "עברית"},
			"ar": {Code: "ar", Flag: "🇸🇦", Name: "العربية"},
			"zh": {Code: "zh", Flag: "🇨🇳", Name: "中文"},
			"ja": {Code: "ja", Flag: "🇯🇵", Name: "日本語"},
			"hi": {Code: "hi", Flag: "🇮🇳", Name: "हिन्दी"},
			"de": {Code: "de", Flag: "🇩🇪", Name: "Deutsch"},
			"th": {Code: "th", Flag: "🇹🇭", Name: "ไทย"},
			"ru": {Code: "ru", Flag: "🇷🇺", Name: "Русский"},          // Rusça
			"pl": {Code: "pl", Flag: "🇵🇱", Name: "Polski"},           // Lehçe
			"fr": {Code: "fr", Flag: "🇫🇷", Name: "Français"},         // Fransızca
			"pt": {Code: "pt", Flag: "🇵🇹", Name: "Português"},        // Portekizce
			"id": {Code: "id", Flag: "🇮🇩", Name: "Bahasa Indonesia"}, // Endonezce
			"bn": {Code: "bn", Flag: "🇧🇩", Name: "বাংলা"},            // Bengalce
		}

		var orientations []payloads.SexualOrientation
		if err := db.Preload("Translations").Find(&orientations).Error; err != nil {
			http.Error(w, "Failed to fetch orientations", http.StatusInternalServerError)
			return
		}

		orientationMap := make(map[string]map[string]string)
		for _, o := range orientations {
			translationMap := make(map[string]string)
			for _, t := range o.Translations {
				translationMap[t.Language] = t.Label
			}
			orientationMap[o.Key] = translationMap
		}

		// 5. InitialData hazırla
		initialData := InitialData{
			Fantasies:          fantasyMap,
			Countries:          countries,
			SexualOrientations: orientationMap,
			Languages:          languages,
			Status:             "ok",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(initialData)
	}
}
