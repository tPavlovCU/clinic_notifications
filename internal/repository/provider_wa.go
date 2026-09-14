package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type WhatsAppProvider struct {
	host       string
	idInstance string
	apiToken   string
	client     *http.Client
}

func InitWhatsAppProvider(host, idInstance, apiToken string, client *http.Client) *WhatsAppProvider {
	return &WhatsAppProvider{
		host:       strings.TrimRight(host, "/"),
		idInstance: idInstance,
		apiToken:   apiToken,
		client:     client,
	}
}

func formatChatID(chatId string) string {
	if strings.HasSuffix(chatId, "@c.us") || strings.HasSuffix(chatId, "@g.us") {
		return chatId
	}

	var sb strings.Builder
	for _, ch := range chatId {
		if ch >= '0' && ch <= '9' {
			sb.WriteRune(ch)
		}
	}
	return sb.String() + "@c.us"
}

func (wa *WhatsAppProvider) SendMessage(ctx context.Context, text string, chatId string) error {
	url := fmt.Sprintf("%s/waInstance%s/sendMessage/%s", wa.host, wa.idInstance, wa.apiToken)

	type greenAPIMessage struct {
		ChatID  string `json:"chatId"`
		Message string `json:"message"`
	}

	msg := greenAPIMessage{
		ChatID:  formatChatID(chatId),
		Message: text,
	}

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("ошибка сериализации json: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonMsg))
	if err != nil {
		return fmt.Errorf("ошибка создания запроса: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := wa.client.Do(req)
	if err != nil {
		return fmt.Errorf("ошибка выполнения http-запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("green api вернул статус код: %d, ответ: %s", resp.StatusCode, string(respBody))
	}

	return nil
}