package main

import (
	"clinic-notifications/internal/config"
	"clinic-notifications/internal/database"
	"clinic-notifications/internal/queue"
	"clinic-notifications/internal/repository"
	"context"
	"fmt"
	"log"
)

var ctx context.Context = context.Background()

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("критическая ошибка при чтении файла env, %v", err)
	}
	fmt.Println("cfg downloaded")

	db, err := database.InitDB(cfg.DBPath)
	if err != nil {
		log.Fatalf("Критическая ошибка при инициализации базы данных, %v", err)
	}
	defer db.Close()
	fmt.Println("db created")

	queue, err := queue.InitRedis(cfg.RedisAddr)
	if err != nil {
		log.Fatalf("Критическая ошибка при инициализации Redis, %v", err)
	}
	defer queue.Close()
	fmt.Println("redis created")
	repo := repository.NewSQLiteRepository(db)
	fmt.Println("repo created")
	repo.SaveLog(ctx, "testAppointmentID", "testPhone", "testTriggerType", "testStatus")
	fmt.Println("log saved")
	fmt.Println(repo.IsAleradySent(ctx, "testAppointmentID", "testTriggerType"))
}
