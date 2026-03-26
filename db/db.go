package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() error {
	var err error
	if DB, err = sql.Open("sqlite3", "api.db"); err != nil {
		return fmt.Errorf("Could not connect to database: %v", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	return createTables()
}

func createTables() error {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER
		)
	`

	if _, err := DB.Exec(createEventsTable); err != nil {
		return fmt.Errorf("Could not create events table: %v", err)
	}

	return nil
}
