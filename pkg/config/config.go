package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type AppConfig struct {
	MainAdminId      []int64 `yaml:"MAIN_ADMIN_ID" env:"MAIN_ADMIN_ID" env-required:"true"`
	TelegramToken    string  `yaml:"TELEGRAM_TOKEN" env:"TELEGRAM_TOKEN" env-required:"true"`
	PollingIsEnabled bool    `yaml:"POLLING_IS_ENABLED" env:"POLLING_IS_ENABLED"`
	PollingInterval  int64   `yaml:"POLLING_INTERVAL" env:"POLLING_INTERVAL" env-required:"true"`
	PollingChatId    int64   `yaml:"POLLING_CHAT_ID" env:"POLLING_CHAT_ID" env-required:"true"`
	KeeneticBaseUrl  string  `yaml:"KEENETIC_BASE_URL" env:"KEENETIC_BASE_URL" env-required:"true"`
	KeeneticUsername string  `yaml:"KEENETIC_USERNAME" env:"KEENETIC_USERNAME" env-required:"true"`
	KeeneticPassword string  `yaml:"KEENETIC_PASSWORD" env:"KEENETIC_PASSWORD" env-required:"true"`
}

func NewConfig() *AppConfig {
	var cfg AppConfig

	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
