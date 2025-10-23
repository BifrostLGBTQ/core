package seeders

import (
	"bifrost/models/user/payloads"

	"gorm.io/gorm"
)

func SeedSexualOrientations(db *gorm.DB) error {
	orientations := []struct {
		Key          string
		Translations map[string]string
	}{
		{"heterosexual", map[string]string{"en": "Heterosexual", "tr": "Heteroseksüel"}},
		{"active_gay", map[string]string{"en": "Active Gay", "tr": "Aktif Gay"}},
		{"passive_gay", map[string]string{"en": "Passive Gay", "tr": "Pasif Gay"}},
		{"versatile_gay", map[string]string{"en": "Versatile Gay", "tr": "Versatil Gay"}},
		{"bisexual_male", map[string]string{"en": "Bisexual Male", "tr": "Biseksüel Erkek"}},
		{"bisexual_female", map[string]string{"en": "Bisexual Female", "tr": "Biseksüel Kadın"}},
		{"lesbian", map[string]string{"en": "Lesbian", "tr": "Lezbiyen"}},
		{"pansexual", map[string]string{"en": "Pansexual", "tr": "Panseksüel"}},
		{"asexual", map[string]string{"en": "Asexual", "tr": "Aseksüel"}},
		{"demisexual", map[string]string{"en": "Demisexual", "tr": "Demiseksüel"}},
		{"queer", map[string]string{"en": "Queer", "tr": "Queer"}},
		{"questioning", map[string]string{"en": "Questioning", "tr": "Sorgulayan"}},
		{"transgender", map[string]string{"en": "Transgender", "tr": "Transgender"}},
		{"other", map[string]string{"en": "Other", "tr": "Diğer"}},
	}

	for i, o := range orientations {
		var orientation payloads.SexualOrientation
		err := db.Where("key = ?", o.Key).First(&orientation).Error
		if err == gorm.ErrRecordNotFound {
			orientation = payloads.SexualOrientation{
				Key:   o.Key,
				Order: i,
			}
			if err := db.Create(&orientation).Error; err != nil {
				return err
			}
		}

		// translations ekle
		for lang, label := range o.Translations {
			var t payloads.SexualOrientationTranslation
			err := db.Where("orientation_id = ? AND language = ?", orientation.ID, lang).First(&t).Error
			if err == gorm.ErrRecordNotFound {
				t = payloads.SexualOrientationTranslation{
					OrientationID: orientation.ID,
					Language:      lang,
					Label:         label,
				}
				if err := db.Create(&t).Error; err != nil {
					return err
				}
			}
		}
	}

	return nil
}
