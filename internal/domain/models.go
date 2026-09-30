package domain

type NotificationTask struct {
	AppointmenId string `json:"appointment_id"`
	ClientPhone  string `json:"client_phone"`
	ClientName   string `json:"client_name"`
	TriggerType  string `json:"trigger_type"`
	MessageText  string `json:"message_text"`
	Date         string `json:"date"`
	CreateDate   string `json:"create_date"`
}

type WebhookDataObject struct {
	Id         int    `json:"id"`
	Date       string `json:"date"`
	CreateDate string `json:"create_date"`
}

type WebhookDataPOST struct {
	Resourse   string            `json:"resourse"`
	ResourceId int               `json:"resource_id"`
	Status     string            `json:"status"`
	Data       WebhookDataObject `json:"data"`
}
