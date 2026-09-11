package database

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
	"os"
	"time"
)

func InitDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть базу данных: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("не удалось подключиться к базе данных: %w", err)
	}
	migrationSql, err := os.ReadFile("migrations/0001_init.sql")

	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать файл миграции: %w", err)
	}

	_, err = db.Exec(string(migrationSql))

	if err != nil {
		return nil, fmt.Errorf("не удалось выполнить миграцию: %w", err, string(migrationSql))
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxIdleTime(5 * time.Minute)
	return db, nil
}
