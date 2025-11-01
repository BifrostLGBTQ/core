package seeders

import (
	"bifrost/models/post/shared"
	"bifrost/models/user/payloads"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedIdentities(db *gorm.DB) error {

	var genderIdentities = []payloads.GenderIdentity{
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Male", "tr": "Erkek"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Female", "tr": "Kadın"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Agender", "tr": "Agender"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Androgynous", "tr": "Androjen"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Aporagender", "tr": "Aporagender"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Bigender", "tr": "Çift cinsiyetli"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Demiboy", "tr": "Yarım erkek"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Demigirl", "tr": "Yarım kız"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Demigender", "tr": "Demicinsiyet"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Genderfluid", "tr": "Cinsiyet akışkan"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Genderflux", "tr": "Genderflux"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Gender neutral", "tr": "Cinsiyetsiz"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Gender questioning", "tr": "Cinsiyet sorgulayan"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Genderqueer", "tr": "Genderqueer"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Graygender", "tr": "Graygender"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Hijara", "tr": "Hijara"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Intergender", "tr": "Intergender"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Intersex", "tr": "Interseks"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Maverique", "tr": "Maverique"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Multigender", "tr": "Çoklu cinsiyet"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Neutrois", "tr": "Neutrois"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Non-binary", "tr": "Non-binary"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Pangender", "tr": "Pangender"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Polygender", "tr": "Poligender"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Transfeminine", "tr": "Transfeminen"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Transmasculine", "tr": "Transmaskülen"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Transneutral", "tr": "Transneutral"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Trigender", "tr": "Trigender"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Two-spirit", "tr": "Two-spirit"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Waria", "tr": "Waria"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Xenogender", "tr": "Xenogender"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Unlabeled", "tr": "Etiketsiz"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "I’d rather not say", "tr": "Belirtmek istemiyorum"}},
	}

	var sexualOrientations = []payloads.SexualOrientation{
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Abrosexual", "tr": "Abroseksüel"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Aceflux", "tr": "Aceflux"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Androsexual", "tr": "Androseksüel"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Aroace", "tr": "Aroace"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Aroflux", "tr": "Aroflux"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Aromantic", "tr": "Aromantik"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Asexual", "tr": "Aseksüel"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "A-Spec", "tr": "A-Spec"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Biromantic", "tr": "Biromantik"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Bisexual", "tr": "Biseksüel"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Ceteroromantic", "tr": "Ceteroromantik"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Ceterosexual", "tr": "Ceteroseksüel"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Demiromantic", "tr": "Demiromantik"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Demisexual", "tr": "Demiseksüel"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Diamoric", "tr": "Diamoric"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Frayromantic", "tr": "Frayromantik"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Gay men", "tr": "Gay erkek"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Gyneromantic", "tr": "Gyneromantik"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Gynesexual", "tr": "Gyneseksüel"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Lesbian", "tr": "Lezbiyen"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Multisexual", "tr": "Multiseksüel"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Omnisexual", "tr": "Omniseksüel"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Pansexual", "tr": "Panseksüel"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Polyamorous", "tr": "Poliamoröz"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Polyromantic", "tr": "Polyromantik"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Polysexual", "tr": "Poliseksüel"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Pomosexual", "tr": "Pomoseksüel"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Queer", "tr": "Queer"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Sapphic", "tr": "Sapphic"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Straight queer", "tr": "Hetero-queer"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Panromantic", "tr": "Panromantik"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Omniromantic", "tr": "Omniromantik"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "I’d rather not say", "tr": "Belirtmek istemiyorum"}},
	}

	var sexRoles = []payloads.SexualRole{
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Top/Active", "tr": "Aktif / Top"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Bottom/Passive", "tr": "Pasif / Bottom"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Versatile/Flexible", "tr": "Versatil / Esnek"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "Switch", "tr": "Switch"}},
		{ID: uuid.New(), Name: shared.LocalizedString{"en": "I’d rather not say", "tr": "Belirtmek istemiyorum"}},
	}
	// --- DB’ye yaz ---
	seedList := []struct {
		slice any
		name  string
	}{
		{genderIdentities, "GenderIdentity"},
		{sexualOrientations, "SexualOrientation"},
		{sexRoles, "SexRole"},
	}

	for _, group := range seedList {
		switch v := group.slice.(type) {
		case []payloads.GenderIdentity:
			for index, item := range v {
				var existing payloads.GenderIdentity
				err := db.Where("name->>'en' = ?", item.Name["en"]).First(&existing).Error
				if err == gorm.ErrRecordNotFound {
					item.DisplayOrder = index
					if err := db.Create(&item).Error; err != nil {
						log.Fatalf("Failed to create %s: %v", group.name, err)
					}
				}
			}

		case []payloads.SexualOrientation:
			for index, item := range v {
				var existing payloads.SexualOrientation
				err := db.Where("name->>'en' = ?", item.Name["en"]).First(&existing).Error
				if err == gorm.ErrRecordNotFound {
					item.DisplayOrder = index
					if err := db.Create(&item).Error; err != nil {
						log.Fatalf("Failed to create %s: %v", group.name, err)
					}
				}
			}

		case []payloads.SexualRole:
			for index, item := range v {
				var existing payloads.SexualRole

				err := db.Where("name->>'en' = ?", item.Name["en"]).First(&existing).Error
				if err == gorm.ErrRecordNotFound {
					item.DisplayOrder = index
					if err := db.Create(&item).Error; err != nil {
						log.Fatalf("Failed to create %s: %v", group.name, err)
					}
				}
			}
		}
	}

	return nil
}
