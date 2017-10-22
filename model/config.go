package model

type TelegramConfig struct {
	Token string `json:"token"`
	DebugEnabled bool `json:"debug_enabled"`
	Offset int `json:"offset"`
	Limit int `json:"limit"`
	Timeout int `json:"timeout"`
}
