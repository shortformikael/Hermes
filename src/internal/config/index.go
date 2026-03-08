package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/shortformikael/Hermes/src/internal/logger"
)

type config struct {
	TelegramToken  string `json:"telegram_token"`
	TelegramChatID string `json:"telegram_chat_id"`
	Port           string `json:"port"`
}

var cfg *config

// Serialize config to JSON

func Init() {
	logger.Info("Initializing config...")
	defer logger.Info("Config initialized successfully")

	// TODO: Add config validation
	// TODO: Add routine to create config.json if it doesn't exist

	configJson, err := os.ReadFile("config.json")
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to read config.json: %v", err))
		os.Exit(1)
	}
	err = json.Unmarshal(configJson, &cfg)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to unmarshal config.json: %v", err))
		os.Exit(1)
	}

	// logger.Info(fmt.Sprintf("Loaded config: %+v", config))
}

func Get() *config {
	return cfg
}
