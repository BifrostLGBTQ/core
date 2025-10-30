package handlers

import (
	"bifrost/models/user/payloads"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

// TranslationMap map[string]Translation

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

type OrientationData struct {
	ID           string            `json:"id"`
	Key          string            `json:"key"`
	Translations map[string]string `json:"translations"`
}

// InitialData dönecek ana struct
type InitialData struct {
	Fantasies []payloads.Fantasy         `json:"fantasies"`
	Countries map[string]CountryResponse `json:"countries"`
	Interests []payloads.Interest        `json:"interests"`

	SexualOrientations []payloads.SexualOrientation `json:"sexual_orientations"` // key -> {lang -> label}
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

		var interests []payloads.Interest
		if err := db.Preload("Items").Find(&interests).Error; err != nil {
			http.Error(w, "Failed to fetch interests", http.StatusInternalServerError)
			return
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

		// 5. InitialData hazırla
		initialData := InitialData{
			Fantasies:          fantasies,
			Countries:          countries,
			Interests:          interests,
			SexualOrientations: orientations,
			Languages:          languages,
			Status:             "ok",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(initialData)
	}
}
