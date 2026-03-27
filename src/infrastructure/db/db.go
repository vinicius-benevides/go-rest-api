package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	Driver       string
	DSN          string
	MaxOpenConns int
	MaxIdleConns int
}

func NewConnection(cfg Config) (*sql.DB, error) {
	connection, err := sql.Open(cfg.Driver, cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("Could not connect to database: %v", err)
	}

	connection.SetMaxOpenConns(cfg.MaxOpenConns)
	connection.SetMaxIdleConns(cfg.MaxIdleConns)

	if err := createTables(connection); err != nil {
		connection.Close()
		return nil, err
	}

	return connection, nil
}

func createTables(db *sql.DB) error {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`

	if _, err := db.Exec(createUsersTable); err != nil {
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

	if _, err := db.Exec(createEventsTable); err != nil {
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

	if _, err := db.Exec(createRegistrationsTable); err != nil {
		return fmt.Errorf("Could not create registrations table: %v", err)
	}

	return nil
}
