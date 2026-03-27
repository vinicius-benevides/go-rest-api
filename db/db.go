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
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`

	if _, err := DB.Exec(createUsersTable); err != nil {
		return fmt.Errorf("Could not create users table: %v", err)
	}

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`

	if _, err := DB.Exec(createEventsTable); err != nil {
		return fmt.Errorf("Could not create events table: %v", err)
	}

	createRegistrationsTable := `
		CREATE TABLE IF NOT EXISTS registrations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_id INTEGER,
			user_id INTEGER,
			FOREIGN KEY (event_id) REFERENCES events(id),
			FOREIGN KEY (user_id) REFERENCES users(id),
			UNIQUE (event_id, user_id)
		)
	`

	if _, err := DB.Exec(createRegistrationsTable); err != nil {
		return fmt.Errorf("Could not create registrations table: %v", err)
	}

	return nil
}
