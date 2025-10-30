package seeders

import (
	"bifrost/models/post/shared"
	"bifrost/models/user/payloads"
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedAttributes(db *gorm.DB) error {

	var HairColors = []payloads.Attribute{
		{ID: uuid.New(), Category: payloads.UserAttributeHairColor, Name: shared.LocalizedString{"en": "Black", "tr": "Siyah"}},
		{ID: uuid.New(), Category: payloads.UserAttributeHairColor, Name: shared.LocalizedString{"en": "Dark Brown", "tr": "Koyu Kahverengi"}},
		{ID: uuid.New(), Category: payloads.UserAttributeHairColor, Name: shared.LocalizedString{"en": "Brown", "tr": "Kahverengi"}},
		{ID: uuid.New(), Category: payloads.UserAttributeHairColor, Name: shared.LocalizedString{"en": "Light Brown", "tr": "Açık Kahverengi"}},
		{ID: uuid.New(), Category: payloads.UserAttributeHairColor, Name: shared.LocalizedString{"en": "Blonde", "tr": "Sarı"}},
		{ID: uuid.New(), Category: payloads.UserAttributeHairColor, Name: shared.LocalizedString{"en": "Red", "tr": "Kızıl"}},
		{ID: uuid.New(), Category: payloads.UserAttributeHairColor, Name: shared.LocalizedString{"en": "Gray", "tr": "Gri"}},
		{ID: uuid.New(), Category: payloads.UserAttributeHairColor, Name: shared.LocalizedString{"en": "White", "tr": "Beyaz"}},
		{ID: uuid.New(), Category: payloads.UserAttributeHairColor, Name: shared.LocalizedString{"en": "Other", "tr": "Diğer"}},
	}

	var EyeColors = []payloads.Attribute{
		{ID: uuid.New(), Category: payloads.UserAttributeEyeColor, Name: shared.LocalizedString{"en": "Brown", "tr": "Kahverengi"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEyeColor, Name: shared.LocalizedString{"en": "Blue", "tr": "Mavi"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEyeColor, Name: shared.LocalizedString{"en": "Green", "tr": "Yeşil"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEyeColor, Name: shared.LocalizedString{"en": "Hazel", "tr": "Ela"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEyeColor, Name: shared.LocalizedString{"en": "Gray", "tr": "Gri"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEyeColor, Name: shared.LocalizedString{"en": "Amber", "tr": "Kehribar"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEyeColor, Name: shared.LocalizedString{"en": "Other", "tr": "Diğer"}},
	}

	var SkinColors = []payloads.Attribute{
		{ID: uuid.New(), Category: payloads.UserAttributeSkinColor, Name: shared.LocalizedString{"en": "Very Fair", "tr": "Çok Açık Ten"}},
		{ID: uuid.New(), Category: payloads.UserAttributeSkinColor, Name: shared.LocalizedString{"en": "Fair", "tr": "Açık Ten"}},
		{ID: uuid.New(), Category: payloads.UserAttributeSkinColor, Name: shared.LocalizedString{"en": "Light", "tr": "Açık Buğday"}},
		{ID: uuid.New(), Category: payloads.UserAttributeSkinColor, Name: shared.LocalizedString{"en": "Medium", "tr": "Buğday"}},
		{ID: uuid.New(), Category: payloads.UserAttributeSkinColor, Name: shared.LocalizedString{"en": "Olive", "tr": "Zeytin Ten"}},
		{ID: uuid.New(), Category: payloads.UserAttributeSkinColor, Name: shared.LocalizedString{"en": "Tan", "tr": "Bronz"}},
		{ID: uuid.New(), Category: payloads.UserAttributeSkinColor, Name: shared.LocalizedString{"en": "Brown", "tr": "Esmer"}},
		{ID: uuid.New(), Category: payloads.UserAttributeSkinColor, Name: shared.LocalizedString{"en": "Dark Brown", "tr": "Koyu Esmer"}},
		{ID: uuid.New(), Category: payloads.UserAttributeSkinColor, Name: shared.LocalizedString{"en": "Black", "tr": "Siyah"}},
	}

	var BodyTypes = []payloads.Attribute{
		{ID: uuid.New(), Category: payloads.UserAttributeBodyType, Name: shared.LocalizedString{"en": "Slim", "tr": "İnce"}},
		{ID: uuid.New(), Category: payloads.UserAttributeBodyType, Name: shared.LocalizedString{"en": "Athletic", "tr": "Atletik"}},
		{ID: uuid.New(), Category: payloads.UserAttributeBodyType, Name: shared.LocalizedString{"en": "Muscular", "tr": "Kaslı"}},
		{ID: uuid.New(), Category: payloads.UserAttributeBodyType, Name: shared.LocalizedString{"en": "Average", "tr": "Orta"}},
		{ID: uuid.New(), Category: payloads.UserAttributeBodyType, Name: shared.LocalizedString{"en": "Chubby", "tr": "Göbekli"}},
		{ID: uuid.New(), Category: payloads.UserAttributeBodyType, Name: shared.LocalizedString{"en": "Heavyset", "tr": "Kilolu"}},
	}

	var Ethnicities = []payloads.Attribute{
		{ID: uuid.New(), Category: payloads.UserAttributeEthnicity, Name: shared.LocalizedString{"en": "White", "tr": "Beyaz"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEthnicity, Name: shared.LocalizedString{"en": "Black", "tr": "Siyah"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEthnicity, Name: shared.LocalizedString{"en": "Hispanic / Latino", "tr": "Hispanik / Latino"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEthnicity, Name: shared.LocalizedString{"en": "Asian", "tr": "Asyalı"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEthnicity, Name: shared.LocalizedString{"en": "East Asian", "tr": "Doğu Asyalı"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEthnicity, Name: shared.LocalizedString{"en": "South Asian", "tr": "Güney Asyalı"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEthnicity, Name: shared.LocalizedString{"en": "Southeast Asian", "tr": "Güneydoğu Asyalı"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEthnicity, Name: shared.LocalizedString{"en": "Middle Eastern / North African", "tr": "Orta Doğulu / Kuzey Afrikalı"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEthnicity, Name: shared.LocalizedString{"en": "Native American / Indigenous", "tr": "Yerli / Kızılderili"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEthnicity, Name: shared.LocalizedString{"en": "Pacific Islander", "tr": "Pasifik Adalı"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEthnicity, Name: shared.LocalizedString{"en": "Mixed", "tr": "Melez"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEthnicity, Name: shared.LocalizedString{"en": "Other", "tr": "Diğer"}},
	}

	var Zodiacs = []payloads.Attribute{
		{ID: uuid.New(), Category: payloads.UserAttributeZodiac, Name: shared.LocalizedString{"en": "Aries", "tr": "Koç"}},
		{ID: uuid.New(), Category: payloads.UserAttributeZodiac, Name: shared.LocalizedString{"en": "Taurus", "tr": "Boğa"}},
		{ID: uuid.New(), Category: payloads.UserAttributeZodiac, Name: shared.LocalizedString{"en": "Gemini", "tr": "İkizler"}},
		{ID: uuid.New(), Category: payloads.UserAttributeZodiac, Name: shared.LocalizedString{"en": "Cancer", "tr": "Yengeç"}},
		{ID: uuid.New(), Category: payloads.UserAttributeZodiac, Name: shared.LocalizedString{"en": "Leo", "tr": "Aslan"}},
		{ID: uuid.New(), Category: payloads.UserAttributeZodiac, Name: shared.LocalizedString{"en": "Virgo", "tr": "Başak"}},
		{ID: uuid.New(), Category: payloads.UserAttributeZodiac, Name: shared.LocalizedString{"en": "Libra", "tr": "Terazi"}},
		{ID: uuid.New(), Category: payloads.UserAttributeZodiac, Name: shared.LocalizedString{"en": "Scorpio", "tr": "Akrep"}},
		{ID: uuid.New(), Category: payloads.UserAttributeZodiac, Name: shared.LocalizedString{"en": "Sagittarius", "tr": "Yay"}},
		{ID: uuid.New(), Category: payloads.UserAttributeZodiac, Name: shared.LocalizedString{"en": "Capricorn", "tr": "Oğlak"}},
		{ID: uuid.New(), Category: payloads.UserAttributeZodiac, Name: shared.LocalizedString{"en": "Aquarius", "tr": "Kova"}},
		{ID: uuid.New(), Category: payloads.UserAttributeZodiac, Name: shared.LocalizedString{"en": "Pisces", "tr": "Balık"}},
	}

	var Disabilities = []payloads.Attribute{
		{ID: uuid.New(), Category: payloads.UserAttributePhysicalDisability, Name: shared.LocalizedString{"en": "Blind", "tr": "Kör"}},
		{ID: uuid.New(), Category: payloads.UserAttributePhysicalDisability, Name: shared.LocalizedString{"en": "Low vision", "tr": "Az görme"}},
		{ID: uuid.New(), Category: payloads.UserAttributePhysicalDisability, Name: shared.LocalizedString{"en": "Deaf", "tr": "Sağır"}},
		{ID: uuid.New(), Category: payloads.UserAttributePhysicalDisability, Name: shared.LocalizedString{"en": "Hard of hearing", "tr": "İşitme zorluğu"}},
		{ID: uuid.New(), Category: payloads.UserAttributePhysicalDisability, Name: shared.LocalizedString{"en": "Wheelchair user", "tr": "Tekerlekli sandalye kullanıcısı"}},
		{ID: uuid.New(), Category: payloads.UserAttributePhysicalDisability, Name: shared.LocalizedString{"en": "Crutches user", "tr": "Koltuk değneği kullanıcısı"}},
		{ID: uuid.New(), Category: payloads.UserAttributePhysicalDisability, Name: shared.LocalizedString{"en": "Amputee (Missing limb)", "tr": "Ampute (Eksik uzuv)"}},
		{ID: uuid.New(), Category: payloads.UserAttributePhysicalDisability, Name: shared.LocalizedString{"en": "Limited arm function", "tr": "Kısıtlı kol fonksiyonu"}},
		{ID: uuid.New(), Category: payloads.UserAttributePhysicalDisability, Name: shared.LocalizedString{"en": "Limited leg function", "tr": "Kısıtlı bacak fonksiyonu"}},
		{ID: uuid.New(), Category: payloads.UserAttributePhysicalDisability, Name: shared.LocalizedString{"en": "Missing hand", "tr": "Eksik el"}},
		{ID: uuid.New(), Category: payloads.UserAttributePhysicalDisability, Name: shared.LocalizedString{"en": "Missing foot", "tr": "Eksik ayak"}},
		{ID: uuid.New(), Category: payloads.UserAttributePhysicalDisability, Name: shared.LocalizedString{"en": "Chronic illness", "tr": "Kronik hastalık"}},
		{ID: uuid.New(), Category: payloads.UserAttributePhysicalDisability, Name: shared.LocalizedString{"en": "Neurological disorder", "tr": "Nörolojik rahatsızlık"}},
		{ID: uuid.New(), Category: payloads.UserAttributePhysicalDisability, Name: shared.LocalizedString{"en": "Respiratory disorder", "tr": "Solunum rahatsızlığı"}},
	}

	var Smoking = []payloads.Attribute{
		{ID: uuid.New(), Category: payloads.UserAttributeSmoking, Name: shared.LocalizedString{"en": "Non-smoker", "tr": "Sigara içmiyor"}},
		{ID: uuid.New(), Category: payloads.UserAttributeSmoking, Name: shared.LocalizedString{"en": "Occasionally", "tr": "Ara sıra"}},
		{ID: uuid.New(), Category: payloads.UserAttributeSmoking, Name: shared.LocalizedString{"en": "Regular smoker", "tr": "Düzenli içiyor"}},
		{ID: uuid.New(), Category: payloads.UserAttributeSmoking, Name: shared.LocalizedString{"en": "I’d rather not say", "tr": "Belirtmek istemiyorum"}},
	}
	var Drinking = []payloads.Attribute{
		{ID: uuid.New(), Category: payloads.UserAttributeDrinking, Name: shared.LocalizedString{"en": "Non-drinker", "tr": "Alkol kullanmıyor"}},
		{ID: uuid.New(), Category: payloads.UserAttributeDrinking, Name: shared.LocalizedString{"en": "Occasionally", "tr": "Ara sıra"}},
		{ID: uuid.New(), Category: payloads.UserAttributeDrinking, Name: shared.LocalizedString{"en": "Social drinker", "tr": "Sosyal içici"}},
		{ID: uuid.New(), Category: payloads.UserAttributeDrinking, Name: shared.LocalizedString{"en": "Regular drinker", "tr": "Düzenli içici"}},
		{ID: uuid.New(), Category: payloads.UserAttributeDrinking, Name: shared.LocalizedString{"en": "I’d rather not say", "tr": "Belirtmek istemiyorum"}},
	}

	var HeightAttributes []payloads.Attribute
	for h := 140; h <= 210; h++ {
		HeightAttributes = append(HeightAttributes, payloads.Attribute{
			ID:       uuid.New(),
			Category: payloads.UserAttributeHeight,
			Name: shared.LocalizedString{
				"en": fmt.Sprintf("%d cm", h),
				"tr": fmt.Sprintf("%d cm", h),
			},
		})
	}

	var WeightAttributes []payloads.Attribute
	for w := 40; w <= 150; w++ {
		WeightAttributes = append(WeightAttributes, payloads.Attribute{
			ID:       uuid.New(),
			Category: payloads.UserAttributeWeight,
			Name: shared.LocalizedString{
				"en": fmt.Sprintf("%d kg", w),
				"tr": fmt.Sprintf("%d kg", w),
			},
		})
	}

	var RelationshipStatuses = []payloads.Attribute{
		{ID: uuid.New(), Category: payloads.UserAttributeRelationshipStatus, Name: shared.LocalizedString{"en": "Single", "tr": "Bekar"}},
		{ID: uuid.New(), Category: payloads.UserAttributeRelationshipStatus, Name: shared.LocalizedString{"en": "Married", "tr": "Evli"}},
		{ID: uuid.New(), Category: payloads.UserAttributeRelationshipStatus, Name: shared.LocalizedString{"en": "Divorced", "tr": "Boşanmış"}},
		{ID: uuid.New(), Category: payloads.UserAttributeRelationshipStatus, Name: shared.LocalizedString{"en": "Widowed", "tr": "Dul"}},
		{ID: uuid.New(), Category: payloads.UserAttributeRelationshipStatus, Name: shared.LocalizedString{"en": "It’s complicated", "tr": "Karmaşık"}},
		{ID: uuid.New(), Category: payloads.UserAttributeRelationshipStatus, Name: shared.LocalizedString{"en": "Open", "tr": "Açık"}},
		{ID: uuid.New(), Category: payloads.UserAttributeRelationshipStatus, Name: shared.LocalizedString{"en": "I’d rather not say", "tr": "Belirtmek istemiyorum"}},
	}

	var KidsPreferences = []payloads.Attribute{
		{ID: uuid.New(), Category: payloads.UserAttributeKidsPreference, Name: shared.LocalizedString{"en": "I’d like them someday", "tr": "Bir gün isterim"}},
		{ID: uuid.New(), Category: payloads.UserAttributeKidsPreference, Name: shared.LocalizedString{"en": "I’d like them soon", "tr": "Yakında isterim"}},
		{ID: uuid.New(), Category: payloads.UserAttributeKidsPreference, Name: shared.LocalizedString{"en": "I don’t want kids", "tr": "Çocuk istemiyorum"}},
		{ID: uuid.New(), Category: payloads.UserAttributeKidsPreference, Name: shared.LocalizedString{"en": "I already have kids", "tr": "Zaten çocuklarım var"}},
		{ID: uuid.New(), Category: payloads.UserAttributeKidsPreference, Name: shared.LocalizedString{"en": "I’d rather not say", "tr": "Belirtmek istemiyorum"}},
	}

	var Pets = []payloads.Attribute{
		{ID: uuid.New(), Category: payloads.UserAttributePets, Name: shared.LocalizedString{"en": "Cat(s)", "tr": "Kedi(ler)"}},
		{ID: uuid.New(), Category: payloads.UserAttributePets, Name: shared.LocalizedString{"en": "Dog(s)", "tr": "Köpek(ler)"}},
		{ID: uuid.New(), Category: payloads.UserAttributePets, Name: shared.LocalizedString{"en": "Both cats and dogs", "tr": "Hem kedi hem köpek"}},
		{ID: uuid.New(), Category: payloads.UserAttributePets, Name: shared.LocalizedString{"en": "Other animals", "tr": "Diğer hayvanlar"}},
		{ID: uuid.New(), Category: payloads.UserAttributePets, Name: shared.LocalizedString{"en": "No pets", "tr": "Hayvan yok"}},
		{ID: uuid.New(), Category: payloads.UserAttributePets, Name: shared.LocalizedString{"en": "I’d rather not say", "tr": "Belirtmek istemiyorum"}},
	}

	var EducationLevels = []payloads.Attribute{
		{ID: uuid.New(), Category: payloads.UserAttributeEducation, Name: shared.LocalizedString{"en": "No formal education", "tr": "Resmî eğitim yok"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEducation, Name: shared.LocalizedString{"en": "Primary school", "tr": "İlkokul"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEducation, Name: shared.LocalizedString{"en": "Middle school", "tr": "Ortaokul"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEducation, Name: shared.LocalizedString{"en": "High school", "tr": "Lise"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEducation, Name: shared.LocalizedString{"en": "Vocational school", "tr": "Meslek lisesi / Meslek yüksekokulu"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEducation, Name: shared.LocalizedString{"en": "Undergraduate degree", "tr": "Lisans"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEducation, Name: shared.LocalizedString{"en": "Graduate degree", "tr": "Yüksek lisans"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEducation, Name: shared.LocalizedString{"en": "Doctorate / PhD", "tr": "Doktora / PhD"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEducation, Name: shared.LocalizedString{"en": "In college", "tr": "Üniversite öğrencisi"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEducation, Name: shared.LocalizedString{"en": "In grad school", "tr": "Yüksek lisans öğrencisi"}},
		{ID: uuid.New(), Category: payloads.UserAttributeEducation, Name: shared.LocalizedString{"en": "I’d rather not say", "tr": "Belirtmek istemiyorum"}},
	}

	var Religions = []payloads.Attribute{
		{ID: uuid.New(), Category: payloads.UserAttributeReligion, Name: shared.LocalizedString{"en": "Agnostic", "tr": "Agnostik"}},
		{ID: uuid.New(), Category: payloads.UserAttributeReligion, Name: shared.LocalizedString{"en": "Atheist", "tr": "Ateist"}},
		{ID: uuid.New(), Category: payloads.UserAttributeReligion, Name: shared.LocalizedString{"en": "Buddhist", "tr": "Budist"}},
		{ID: uuid.New(), Category: payloads.UserAttributeReligion, Name: shared.LocalizedString{"en": "Catholic", "tr": "Katolik"}},
		{ID: uuid.New(), Category: payloads.UserAttributeReligion, Name: shared.LocalizedString{"en": "Christian", "tr": "Hristiyan"}},
		{ID: uuid.New(), Category: payloads.UserAttributeReligion, Name: shared.LocalizedString{"en": "Hindu", "tr": "Hindu"}},
		{ID: uuid.New(), Category: payloads.UserAttributeReligion, Name: shared.LocalizedString{"en": "Jain", "tr": "Jain"}},
		{ID: uuid.New(), Category: payloads.UserAttributeReligion, Name: shared.LocalizedString{"en": "Jewish", "tr": "Yahudi"}},
		{ID: uuid.New(), Category: payloads.UserAttributeReligion, Name: shared.LocalizedString{"en": "Mormon", "tr": "Mormon"}},
		{ID: uuid.New(), Category: payloads.UserAttributeReligion, Name: shared.LocalizedString{"en": "Muslim", "tr": "Müslüman"}},
		{ID: uuid.New(), Category: payloads.UserAttributeReligion, Name: shared.LocalizedString{"en": "Zoroastrian", "tr": "Zerdüşt"}},
		{ID: uuid.New(), Category: payloads.UserAttributeReligion, Name: shared.LocalizedString{"en": "Sikh", "tr": "Sih"}},
		{ID: uuid.New(), Category: payloads.UserAttributeReligion, Name: shared.LocalizedString{"en": "Spiritual", "tr": "Spiritüel"}},
		{ID: uuid.New(), Category: payloads.UserAttributeReligion, Name: shared.LocalizedString{"en": "Baháʼí", "tr": "Bahai"}},
		{ID: uuid.New(), Category: payloads.UserAttributeReligion, Name: shared.LocalizedString{"en": "Shinto", "tr": "Şinto"}},
		{ID: uuid.New(), Category: payloads.UserAttributeReligion, Name: shared.LocalizedString{"en": "Taoism", "tr": "Taoizm"}},
		{ID: uuid.New(), Category: payloads.UserAttributeReligion, Name: shared.LocalizedString{"en": "Confucianism", "tr": "Konfüçyüsçülük"}},
		{ID: uuid.New(), Category: payloads.UserAttributeReligion, Name: shared.LocalizedString{"en": "Animism", "tr": "Animizm"}},
		{ID: uuid.New(), Category: payloads.UserAttributeReligion, Name: shared.LocalizedString{"en": "Pagan", "tr": "Pagan"}},
		{ID: uuid.New(), Category: payloads.UserAttributeReligion, Name: shared.LocalizedString{"en": "Rastafarian", "tr": "Rastafari"}},
		{ID: uuid.New(), Category: payloads.UserAttributeReligion, Name: shared.LocalizedString{"en": "Indigenous", "tr": "Yerli inançlar"}},
		{ID: uuid.New(), Category: payloads.UserAttributeReligion, Name: shared.LocalizedString{"en": "Other", "tr": "Diğer"}},
		{ID: uuid.New(), Category: payloads.UserAttributeReligion, Name: shared.LocalizedString{"en": "I’d rather not say", "tr": "Belirtmek istemiyorum"}},
	}

	allAttributes := [][]payloads.Attribute{HeightAttributes, WeightAttributes, HairColors, EyeColors, SkinColors, BodyTypes, Ethnicities, Zodiacs, RelationshipStatuses, KidsPreferences, EducationLevels, Religions, Disabilities, Pets, Smoking, Drinking}

	for _, attrs := range allAttributes {
		for index, attr := range attrs {
			var existing payloads.Attribute
			err := db.Where("category = ? AND name->>'en' = ?", attr.Category, attr.Name["en"]).First(&existing).Error
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					attr.DisplayOrder = index
					if err := db.Create(&attr).Error; err != nil {
						log.Fatalf("Failed to create attribute %v: %v", attr.Name, err)
					}
				} else {
					log.Fatalf("DB error: %v", err)
				}
			}
		}
	}

	return nil
}
