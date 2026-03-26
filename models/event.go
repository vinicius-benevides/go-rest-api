package models

import (
	"fmt"
	"time"

	"github.com/vinicius-benevides/go-rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	query := `
		INSERT INTO events(name, description, location, dateTime, user_id) 
		VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("Could not prepare statement for saving event: %v", err)
	}

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return fmt.Errorf("Could not execute statement for saving event: %v", err)
	}
	defer stmt.Close()

	e.ID, err = result.LastInsertId()
	if err != nil {
		return fmt.Errorf("Could not get last inserted id for saved event: %v", err)
	}

	return nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Could not get all events: %v", err)
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err = rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserID,
		)
		if err != nil {
			return nil, fmt.Errorf("Could not get all events: %v", err)
		}

		events = append(events, event)
	}

	return events, nil
}
