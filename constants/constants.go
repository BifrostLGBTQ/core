package constants

import (
	"bifrost/models"
)

const (
	GenderMale        models.GenderIdentity = "male"
	GenderFemale      models.GenderIdentity = "female"
	GenderTransMale   models.GenderIdentity = "trans_male"
	GenderTransFemale models.GenderIdentity = "trans_female"
	GenderNonBinary   models.GenderIdentity = "non_binary"
	GenderOther       models.GenderIdentity = "other"
	GenderAgender     models.GenderIdentity = "agender"
	GenderBigender    models.GenderIdentity = "bigender"
	GenderGenderfluid models.GenderIdentity = "genderfluid"
	GenderDemiboy     models.GenderIdentity = "demiboy"
	GenderDemigirl    models.GenderIdentity = "demigirl"
	GenderTwoSpirit   models.GenderIdentity = "two_spirit"
	GenderNeutrois    models.GenderIdentity = "neutrois"
	GenderIntersex    models.GenderIdentity = "intersex"
	GenderQuestioning models.GenderIdentity = "questioning"
	GenderX           models.GenderIdentity = "x"

	OrientationGay         models.SexualOrientation = "gay"
	OrientationLesbian     models.SexualOrientation = "lesbian"
	OrientationBi          models.SexualOrientation = "bisexual"
	OrientationPan         models.SexualOrientation = "pansexual"
	OrientationAsexual     models.SexualOrientation = "asexual"
	OrientationStraight    models.SexualOrientation = "straight"
	OrientationOther       models.SexualOrientation = "other"
	OrientationDemisexual  models.SexualOrientation = "demisexual"
	OrientationQueer       models.SexualOrientation = "queer"
	OrientationGraysexual  models.SexualOrientation = "graysexual"
	OrientationQuestioning models.SexualOrientation = "questioning"
	OrientationSapiosexual models.SexualOrientation = "sapiosexual"
	OrientationAndrosexual models.SexualOrientation = "androsexual"
	OrientationGynosexual  models.SexualOrientation = "gynosexual"

	// Relationship Status
	StatusSingle       models.RelationshipStatus = "single"
	StatusRelationship models.RelationshipStatus = "relationship"
	StatusMarriage     models.RelationshipStatus = "marriage"
	StatusPartnership  models.RelationshipStatus = "partnership"
	StatusInBetween    models.RelationshipStatus = "inbetween"
	StatusIDK          models.RelationshipStatus = "idk"
	StatusDivorced     models.RelationshipStatus = "divorced"
	StatusWidowed      models.RelationshipStatus = "widowed"
	StatusSeparated    models.RelationshipStatus = "separated"
	StatusOpen         models.RelationshipStatus = "open"
	StatusEngaged      models.RelationshipStatus = "engaged"

	// BDSM Interest (İlgi)
	BDSMInterestYes         models.BDSMInterest = "yes"
	BDSMInterestNo          models.BDSMInterest = "no"
	BDSMInterestCurious     models.BDSMInterest = "curious"
	BDSMInterestExperienced models.BDSMInterest = "experienced"
	BDSMInterestOther       models.BDSMInterest = "other"

	// BDSM Roles (Roller)
	BDSMRoleDominant   models.BDSMRole = "dominant"
	BDSMRoleSubmissive models.BDSMRole = "submissive"
	BDSMRoleSwitch     models.BDSMRole = "switch"
	BDSMRoleMaster     models.BDSMRole = "master"
	BDSMRoleSlave      models.BDSMRole = "slave"
	BDSMRoleTop        models.BDSMRole = "top"
	BDSMRoleBottom     models.BDSMRole = "bottom"
	BDSMRoleNone       models.BDSMRole = "none"

	RoleTop           models.SexRole = "top"
	RoleBottom        models.SexRole = "bottom"
	RoleVerse         models.SexRole = "verse"
	RoleSide          models.SexRole = "side"
	RoleNone          models.SexRole = "none"
	RoleVersTop       models.SexRole = "vers-top"
	RoleVersBottom    models.SexRole = "vers-bottom"
	RoleServiceTop    models.SexRole = "service-top"
	RoleServiceBottom models.SexRole = "service-bottom"
	RoleDominant      models.SexRole = "dominant"
	RoleSubmissive    models.SexRole = "submissive"
	RoleSwitch        models.SexRole = "switch"
	RoleBrat          models.SexRole = "brat"
	RoleObserver      models.SexRole = "observer"
	RoleExhibitionist models.SexRole = "exhibitionist"

	ZodiacAries       models.ZodiacSign = "aries"       // Koç
	ZodiacTaurus      models.ZodiacSign = "taurus"      // Boğa
	ZodiacGemini      models.ZodiacSign = "gemini"      // İkizler
	ZodiacCancer      models.ZodiacSign = "cancer"      // Yengeç
	ZodiacLeo         models.ZodiacSign = "leo"         // Aslan
	ZodiacVirgo       models.ZodiacSign = "virgo"       // Başak
	ZodiacLibra       models.ZodiacSign = "libra"       // Terazi
	ZodiacScorpio     models.ZodiacSign = "scorpio"     // Akrep
	ZodiacSagittarius models.ZodiacSign = "sagittarius" // Yay
	ZodiacCapricorn   models.ZodiacSign = "capricorn"   // Oğlak
	ZodiacAquarius    models.ZodiacSign = "aquarius"    // Kova
	ZodiacPisces      models.ZodiacSign = "pisces"      // Balık
	ZodiacUnknown     models.ZodiacSign = "unknown"     // Bilinmiyor / Belirtilmemiş

	SmokingNever        models.SmokingHabit = "never"
	SmokingOccasionally models.SmokingHabit = "occasionally"
	SmokingRegularly    models.SmokingHabit = "regularly"
	SmokingTryingToQuit models.SmokingHabit = "trying_to_quit"
	SmokingOther        models.SmokingHabit = "other"

	DrinkingNever        models.DrinkingHabit = "never"
	DrinkingOccasionally models.DrinkingHabit = "occasionally"
	DrinkingRegularly    models.DrinkingHabit = "regularly"
	DrinkingTryingToQuit models.DrinkingHabit = "trying_to_quit"
	DrinkingOther        models.DrinkingHabit = "other"

	UserRoleUser       models.UserRole = "user"
	UserRoleModerator  models.UserRole = "moderator"
	UserRoleAdmin      models.UserRole = "admin"
	UserRoleSuperAdmin models.UserRole = "super_admin"
	UserRoleBanned     models.UserRole = "banned"
	UserRoleDeleted    models.UserRole = "deleted"
	UserRolePending    models.UserRole = "pending"
	UserRoleVerified   models.UserRole = "verified"
	UserRoleUnverified models.UserRole = "unverified"
)
