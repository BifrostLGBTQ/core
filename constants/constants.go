package constants

type FollowStatus string
type GenderIdentity string
type SexRole string
type UserRole string
type RelationshipStatus string
type BDSMInterest string
type BDSMRole string

type ZodiacSign string

type SmokingHabit string
type DrinkingHabit string

type TravelPurpose string

const (
	GenderMale        GenderIdentity = "male"
	GenderFemale      GenderIdentity = "female"
	GenderTransMale   GenderIdentity = "trans_male"
	GenderTransFemale GenderIdentity = "trans_female"
	GenderNonBinary   GenderIdentity = "non_binary"
	GenderOther       GenderIdentity = "other"
	GenderAgender     GenderIdentity = "agender"
	GenderBigender    GenderIdentity = "bigender"
	GenderGenderfluid GenderIdentity = "genderfluid"
	GenderDemiboy     GenderIdentity = "demiboy"
	GenderDemigirl    GenderIdentity = "demigirl"
	GenderTwoSpirit   GenderIdentity = "two_spirit"
	GenderNeutrois    GenderIdentity = "neutrois"
	GenderIntersex    GenderIdentity = "intersex"
	GenderQuestioning GenderIdentity = "questioning"
	GenderX           GenderIdentity = "x"

	/*
		OrientationGay         SexualOrientation = "gay"
		OrientationLesbian     SexualOrientation = "lesbian"
		OrientationBi          SexualOrientation = "bisexual"
		OrientationPan         SexualOrientation = "pansexual"
		OrientationAsexual     SexualOrientation = "asexual"
		OrientationStraight    SexualOrientation = "straight"
		OrientationOther       SexualOrientation = "other"
		OrientationDemisexual  SexualOrientation = "demisexual"
		OrientationQueer       SexualOrientation = "queer"
		OrientationGraysexual  SexualOrientation = "graysexual"
		OrientationQuestioning SexualOrientation = "questioning"
		OrientationSapiosexual SexualOrientation = "sapiosexual"
		OrientationAndrosexual SexualOrientation = "androsexual"
		OrientationGynosexual  SexualOrientation = "gynosexual"
	*/
	// Relationship Status
	StatusSingle       RelationshipStatus = "single"
	StatusRelationship RelationshipStatus = "relationship"
	StatusMarriage     RelationshipStatus = "marriage"
	StatusPartnership  RelationshipStatus = "partnership"
	StatusInBetween    RelationshipStatus = "inbetween"
	StatusIDK          RelationshipStatus = "idk"
	StatusDivorced     RelationshipStatus = "divorced"
	StatusWidowed      RelationshipStatus = "widowed"
	StatusSeparated    RelationshipStatus = "separated"
	StatusOpen         RelationshipStatus = "open"
	StatusEngaged      RelationshipStatus = "engaged"

	// BDSM Interest (İlgi)
	BDSMInterestYes         BDSMInterest = "yes"
	BDSMInterestNo          BDSMInterest = "no"
	BDSMInterestCurious     BDSMInterest = "curious"
	BDSMInterestExperienced BDSMInterest = "experienced"
	BDSMInterestOther       BDSMInterest = "other"

	// BDSM Roles (Roller)
	BDSMRoleDominant   BDSMRole = "dominant"
	BDSMRoleSubmissive BDSMRole = "submissive"
	BDSMRoleSwitch     BDSMRole = "switch"
	BDSMRoleMaster     BDSMRole = "master"
	BDSMRoleSlave      BDSMRole = "slave"
	BDSMRoleTop        BDSMRole = "top"
	BDSMRoleBottom     BDSMRole = "bottom"
	BDSMRoleNone       BDSMRole = "none"

	RoleTop           SexRole = "top"
	RoleBottom        SexRole = "bottom"
	RoleVerse         SexRole = "verse"
	RoleSide          SexRole = "side"
	RoleNone          SexRole = "none"
	RoleVersTop       SexRole = "vers-top"
	RoleVersBottom    SexRole = "vers-bottom"
	RoleServiceTop    SexRole = "service-top"
	RoleServiceBottom SexRole = "service-bottom"
	RoleDominant      SexRole = "dominant"
	RoleSubmissive    SexRole = "submissive"
	RoleSwitch        SexRole = "switch"
	RoleBrat          SexRole = "brat"
	RoleObserver      SexRole = "observer"
	RoleExhibitionist SexRole = "exhibitionist"

	ZodiacAries       ZodiacSign = "aries"       // Koç
	ZodiacTaurus      ZodiacSign = "taurus"      // Boğa
	ZodiacGemini      ZodiacSign = "gemini"      // İkizler
	ZodiacCancer      ZodiacSign = "cancer"      // Yengeç
	ZodiacLeo         ZodiacSign = "leo"         // Aslan
	ZodiacVirgo       ZodiacSign = "virgo"       // Başak
	ZodiacLibra       ZodiacSign = "libra"       // Terazi
	ZodiacScorpio     ZodiacSign = "scorpio"     // Akrep
	ZodiacSagittarius ZodiacSign = "sagittarius" // Yay
	ZodiacCapricorn   ZodiacSign = "capricorn"   // Oğlak
	ZodiacAquarius    ZodiacSign = "aquarius"    // Kova
	ZodiacPisces      ZodiacSign = "pisces"      // Balık
	ZodiacUnknown     ZodiacSign = "unknown"     // Bilinmiyor / Belirtilmemiş

	SmokingNever        SmokingHabit = "never"
	SmokingOccasionally SmokingHabit = "occasionally"
	SmokingRegularly    SmokingHabit = "regularly"
	SmokingTryingToQuit SmokingHabit = "trying_to_quit"
	SmokingOther        SmokingHabit = "other"

	DrinkingNever        DrinkingHabit = "never"
	DrinkingOccasionally DrinkingHabit = "occasionally"
	DrinkingRegularly    DrinkingHabit = "regularly"
	DrinkingTryingToQuit DrinkingHabit = "trying_to_quit"
	DrinkingOther        DrinkingHabit = "other"

	UserRoleUser       UserRole = "user"
	UserRoleModerator  UserRole = "moderator"
	UserRoleAdmin      UserRole = "admin"
	UserRoleSuperAdmin UserRole = "super_admin"
	UserRoleBanned     UserRole = "banned"
	UserRoleDeleted    UserRole = "deleted"
	UserRolePending    UserRole = "pending"
	UserRoleVerified   UserRole = "verified"
	UserRoleUnverified UserRole = "unverified"
)
