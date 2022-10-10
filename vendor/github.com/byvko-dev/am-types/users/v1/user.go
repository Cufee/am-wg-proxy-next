package users

type User struct {
	DefaultPID int    `json:"player_id"`
	Locale     string `json:"locale"`

	Premium  bool `json:"premium"`
	Verified bool `json:"verified"`

	CustomBgURL string `json:"bg_url"`

	Banned       bool   `json:"banned"`
	BanReason    string `json:"ban_reason,omitempty"`
	BanNotified  bool   `json:"ban_notified,omitempty"`
	ShadowBanned bool   `json:"shadow_banned,omitempty"`
}
