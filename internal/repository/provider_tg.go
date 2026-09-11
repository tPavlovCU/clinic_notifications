package repository

import (
	"net/http"
)

type TelegramProvider struct {
	token  string
	client *http.Client
}

func InitTelegramProvider(token string)
