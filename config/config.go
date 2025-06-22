package config

import (
	"os"
)

// Config holds all configuration for the application
type Config struct {
	Telegram TelegramConfig
}

// TelegramConfig holds Telegram bot configuration
type TelegramConfig struct {
	BotToken string
	ChatID   string
}

// Load loads configuration from environment variables
func Load() *Config {
	return &Config{
		Telegram: TelegramConfig{
			BotToken: getEnv("TELEGRAM_BOT_TOKEN", ""),
			ChatID:   getEnv("TELEGRAM_CHAT_ID", ""),
		},
	}
}

// getEnv gets an environment variable with a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
