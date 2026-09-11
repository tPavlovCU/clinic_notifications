package domain

import (
	"context"
	"fmt"
)

type DBRepository interface {
	IsAlreadySent(ctx context.Context, appointmentId, triggerType string) (bool, error)
	SaveLog(ctx context.Context, appointmentID, phone, triggerType, status string) error
}

type NotificationSendler interface {
	SendMessage(ctx context.Context, text string, chatId string) error
}

type NotificationService struct {
	repo    DBRepository
	sendler NotificationSendler
}

func NewNotificationService(repo DBRepository, sendler NotificationSendler) *NotificationService {
	res := NotificationService{
		repo:    repo,
		sendler: sendler,
	}
	return &res
}

func (ns *NotificationService) HandleBooking(ctx context.Context, appointmentId, phone, clientName, chatId string) error {
	sent, err := ns.repo.IsAlreadySent(ctx, appointmentId, "booking")
	if err != nil {
		return err
	}
	if sent == true {
		return nil
	}
	msg := fmt.Sprintf("Здравствуйте, %s! Вы успешно записаны в клинику. ID записи: %s", clientName, appointmentId)
	err = ns.sendler.SendMessage(ctx, msg, chatId)
	if err != nil {
		saveErr := ns.repo.SaveLog(ctx, appointmentId, phone, "booking", "failed")
		if saveErr != nil {
			fmt.Println("Ошибка при записи лога!", saveErr)
		}
		return err
	}
	saveErr := ns.repo.SaveLog(ctx, appointmentId, phone, "booking", "success")
	if saveErr != nil {
		fmt.Println("Ошибка при записи лога!", saveErr)
	}

	return nil

}
