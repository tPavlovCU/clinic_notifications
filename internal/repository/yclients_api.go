package repository

import (
	"context"
	"encoding/json"
)

type YclientsAPI struct {
	token string
}

type YclientsClient struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type YclientsAppointment struct {
	ID     int            `json:"id"`
	Client YclientsClient `json:"client"`
}

type YclientsResponse struct {
	Data []YclientsAppointment `json:"data"`
}

func NewClientAPI(token string) *YclientsAPI {
	return &YclientsAPI{token}
}

func (y *YclientsAPI) GetRecentAppointments(ctx context.Context) ([]YclientsAppointment, error) {
	mockJSON := `{
    	"data": [
        	{"id": 1001, "client": {"name": "Алексей", "phone": "79992223344"}},
        	{"id": 1002, "client": {"name": "Мария", "phone": "79995556677"}}
    	]
	}`

	var resp YclientsResponse

	err := json.Unmarshal([]byte(mockJSON), &resp)
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}
