package database

import (
	"database/sql"
	_ "modernc.org/sqlite"
	"fmt"
	"os"
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
		return nil, fmt.Errorf("не удалось выполнить миграцию: %w", err)
	}
	return db, nil
}