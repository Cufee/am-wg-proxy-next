package types

type AchievementDetails struct {
	ID          string `json:"achievement_id" bson:"achievement_id"`
	Name        string `json:"realm" bson:"name"`
	Section     string `json:"section" bson:"section"`
	ImageURL    string `json:"image" bson:"image"`
	Description string `json:"description" bson:"description"`
	Condition   string `json:"condition" bson:"condition"`
}

const (
	LangEN   = "en"    // English
	LangRU   = "ru"    // Russian
	LangPL   = "pl"    // Polish
	LangDE   = "de"    // German
	LangFR   = "fr"    // French
	LangES   = "es"    // Spanish
	LangTR   = "tr"    // Turkish
	LangCS   = "cs"    // Czech
	LangTH   = "th"    // Thai
	LangKO   = "ko"    // Korean
	LangVI   = "vi"    // Vietnamese
	LangZhCH = "zh-cn" // Simplified Chinese
	LangZhTW = "zh-tw" // Traditional Chinese
)

var AllLanguages = []string{LangEN, LangRU, LangPL, LangDE, LangFR, LangES, LangTR, LangCS, LangTH, LangKO, LangVI, LangZhCH, LangZhTW}

func GetLocale(locale string) string {
	for _, l := range AllLanguages {
		if l == locale {
			return l
		}
	}
	return LangEN
}

type VehicleDetails struct {
	TankID    int    `json:"tank_id,omitempty" bson:"tank_id,omitempty"`
	Name      string `json:"name,omitempty" bson:"name,omitempty"`
	Nation    string `json:"nation,omitempty" bson:"nation,omitempty"`
	Tier      int    `json:"tier,omitempty" bson:"tier,omitempty"`
	Type      string `json:"type,omitempty" bson:"type,omitempty"`
	IsPremium bool   `json:"is_premium,omitempty" bson:"is_premium,omitempty"`
}
