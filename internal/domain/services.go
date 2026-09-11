package domain

import (
	"context"
)

type DBRepository interface {
	IsAlreadySent(ctx context.Context, appointmentId, triggerType string) (bool, error)
	SaveLog(ctx context.Context, appointmentID, phone, triggerType, status string) error
}

type NotificationSendler interface {
	SendMessage(ctx context.Context, text string, chatId string) error
}

type NotificationSerice struct {
	repo    DBRepository
	sendler NotificationSendler
}
