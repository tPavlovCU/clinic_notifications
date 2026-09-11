package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type TelegramProvider struct {
	token  string
	client *http.Client
}

func InitTelegramProvider(token string, client *http.Client) *TelegramProvider {
	tg := TelegramProvider{
		token:  token,
		client: client,
	}
	return &tg
}

func (tg *TelegramProvider) SendMessage(ctx context.Context, text string, chatId string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage", tg.token)
	type message struct {
		Text   string `json:"text"`
		ChatId string `json:"chat_id"`
	}
	msg := message{
		Text:   text,
		ChatId: chatId,
	}
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonMsg))
	if err != nil {
		return err
	}
	req.Header.Set("Content-type", "application/json")
	resp, err := tg.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("телеграм вернул статус код: %v", resp.StatusCode)
	}
	return nil
}
