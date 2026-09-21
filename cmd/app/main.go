package main

//main file
import (
	"clinic-notifications/internal/config"
	"clinic-notifications/internal/database"
	//"clinic-notifications/internal/queue"
	"clinic-notifications/internal/domain"
	"clinic-notifications/internal/repository"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var ctx context.Context = context.Background()

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

	// queue, err := queue.InitRedis(cfg.RedisAddr)
	// if err != nil {
	// 	log.Fatalf("Критическая ошибка при инициализации Redis, %v", err)
	// }
	// defer queue.Close()
	repo := repository.NewSQLiteRepository(db)
	customHTTPClient := &http.Client{
		Timeout: time.Second * 5,
	}

	// .............................TELEGRAM PART
	// tgProvider := repository.InitTelegramProvider(cfg.TelegramBotToken, customHTTPClient)
	// serviceTG := domain.NewNotificationService(repo, tgProvider)
	// err = serviceTG.HandleBooking(ctx, "appt_999", "79991112233", "Иван", "фейковый_chat_id")

	// yclients := repository.NewClientAPI(cfg.YclientsToken)

	// ticker := time.NewTicker(10 * time.Second)
	// defer ticker.Stop()
	// fmt.Println("Сервис полностью запущен и ожидает новые записи...")
	// for range ticker.C {
	// 	appointments, err := yclients.GetRecentAppointments(ctx)
	// 	if err != nil {
	// 		log.Println("Ошибка при получении данных")
	// 	}
	// 	for _, appt := range appointments {
	// 		idStr := strconv.Itoa(appt.ID)

	// 		err := serviceTG.HandleBooking(ctx, idStr, appt.Client.Phone, appt.Client.Name, "тестовый_chat_id")
	// 		if err != nil {
	// 			log.Println("Ошибка в booking")
	// 		}

	// 	}
	// }

	// .............................WHATSAPP PART
	waProvider := repository.InitWhatsAppProvider(
		cfg.GreenApiHost,
		cfg.GreenApiIDInstance,
		cfg.GreenApiToken,
		customHTTPClient,
	)

	serviceWA := domain.NewNotificationService(repo, waProvider)

	yclients := repository.NewClientAPI(cfg.YclientsToken)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	fmt.Println("Сервис полностью запущен и ожидает новые записи...")
	for range ticker.C {
		appointments, err := yclients.GetRecentAppointments(ctx)
		if err != nil {
			log.Println("Ошибка при получении данных", err)
		}
		for _, appt := range appointments {
			idStr := strconv.Itoa(appt.ID)

			err := serviceWA.HandleBooking(ctx, idStr, appt.Client.Phone, appt.Client.Name, appt.Client.Phone)
			if err != nil {
				log.Println("Ошибка в booking", err)
			}
			fmt.Println("all ok", err)
			log.Printf("Ошибка при отправке в WhatsApp для записи %s: %v\n", idStr, err)
		}
	}
}
