package config

import (
	"os"

	"github.com/shortformikael/Hermes/src/internal/logger"
)

type Config struct {
	TelegramToken  string
	TelegramChatID string
	Port           string
}

var config *Config

// Serialize config to JSON

func Init() {
	logger.Info("Initializing config...")
	defer logger.Info("Config initialized successfully")
	config = &Config{
		TelegramToken:  os.Getenv("TELEGRAM_TOKEN"),
		TelegramChatID: os.Getenv("TELEGRAM_CHAT_ID"),
		Port:           os.Getenv("PORT"),
	}
}

func GetConfig() *Config {
	return config
}
