package main

import (
	"clinic-notifications/internal/config"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("критическая ошибка при чтении файла env, %v", err)
	}

	fmt.Println(cfg)
	//just comment
}
