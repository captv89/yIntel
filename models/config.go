package models

type Config struct {
	TelegramToken string  `yaml:"telegram_token"`
	AllowedIds    []int64 `yaml:"allowed_ids"`
}
