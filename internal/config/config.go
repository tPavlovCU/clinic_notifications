package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	ServerPort       string
	DBPath           string
	RedisAddr        string
	YclientsToken    string
	WhatsAppApiURL   string
	TelegramBotToken string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Файл env не найден")
	}

	cfg := Config{
		getEnv("SERVER_PORT", "8080"),
		getEnv("DB_PATH", "clinic.db"),
		getEnv("REDIS_ADDR", "localhost:6379"),
		os.Getenv("YCLIENTS_API_TOKEN"),
		os.Getenv("WHATSAPP_API_URL"),
		os.Getenv("TELEGRAM_BOT_TOKEN"),
	}

	return &cfg, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
