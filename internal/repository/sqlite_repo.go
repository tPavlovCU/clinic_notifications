package repository

import (
	"context"
	"database/sql"
	_ "modernc.org/sqlite"
)

type SQLRepo struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLRepo {
	rep := &SQLRepo{db}
	return rep
}

func (r *SQLRepo) IsAlreadySent(ctx context.Context, appointmentId, triggerType string) (bool, error) {
	queue := "SELECT COUNT(*) FROM logs WHERE appointment_id = ? AND trigger_type = ?" //может сюда стоит добавить еще STATUS = success? типо что сообщение отправленое успешно
	result := r.db.QueryRowContext(ctx, queue, appointmentId, triggerType)

	var count int
	err := result.Scan(&count)
	if err != nil {
		return true, nil
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (r *SQLRepo) SaveLog(ctx context.Context, appointmentID, phone, triggerType, status string) error {
	queue := "INSERT INTO logs (appointment_id, client_phone, trigger_type, status) VALUES (?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, queue, appointmentID, phone, triggerType, status)
	if err != nil {
		return err
	}
	return nil
}
