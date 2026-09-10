package domain

import ()

type NotificationTask struct {
	AppointmenId string `json:"appointment_id"`
	ClientPhone  string `json:"client_phone"`
	ClientName   string `json:"client_name"`
	TriggerType  string `json:"trigger_type"`
	MessageText  string `json:"message_text"`
}
