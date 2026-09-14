package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	ServerPort			string
	DBPath				string
	RedisAddr			string
	YclientsToken		string
	TelegramBotToken	string
	GreenApiHost		string
	GreenApiIDInstance	string
	GreenApiToken		string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Файл env не найден")
	}

	cfg := Config{
		ServerPort:			getEnv("SERVER_PORT", "8080"),
		DBPath:				getEnv("DB_PATH", "clinic.db"),
		RedisAddr:			getEnv("REDIS_ADDR", "localhost:6379"),
		YclientsToken:		os.Getenv("YCLIENTS_API_TOKEN"),
		TelegramBotToken:	os.Getenv("TELEGRAM_BOT_TOKEN"),
		GreenApiHost:		getEnv("GREEN_API_HOST", "https://api.green-api.com"),
		GreenApiIDInstance:	os.Getenv("GREEN_API_ID_INSTANCE"),
		GreenApiToken:		os.Getenv("GREEN_API_TOKEN_INSTANCE"),
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
