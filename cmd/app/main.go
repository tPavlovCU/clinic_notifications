package main

import (
	"clinic-notifications/internal/config"
	"clinic-notifications/internal/database"
	"clinic-notifications/internal/queue"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("критическая ошибка при чтении файла env, %v", err)
	}

	db, err := database.InitDB(cfg.DBPath)
	if err != nil {
		log.Fatalf("Критическая ошибка при инициализации базы данных, %v", err)
	}
	defer db.Close()

	queue, err := queue.InitRedis(cfg.RedisAddr)
	if err != nil {
		log.Fatalf("Критическая ошибка при инициализации Redis, %v", err)
	}
	defer queue.Close()

}
