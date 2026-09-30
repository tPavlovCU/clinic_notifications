package repository

import (
	"clinic-notifications/internal/domain"
	"encoding/json"
	"net/http"
)

type WebhookHandler struct {
	service *domain.NotificationService
}

func NewWebhookHandler(service *domain.NotificationService) *WebhookHandler {
	return &WebhookHandler{service: service}
}

func (h *WebhookHandler) RegisterWebhook(mux *http.ServeMux) {
	mux.HandleFunc("/webhook", h.HandleWebhook)
}

func (h *WebhookHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	var Data domain.WebhookDataPOST

	json.NewDecoder(r.Body).Decode(&Data)

	switch Data.Status {
	case "create":
		a := 1 + 1
		_ = a
	case "update":
		a := 1 + 1
		_ = a
	case "delete":
		a := 1 + 1
		_ = a
	} //"create" "update" "delete"
}
