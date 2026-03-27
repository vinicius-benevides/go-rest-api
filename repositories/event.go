package repositories

import (
	"database/sql"
	"fmt"

	"github.com/vinicius-benevides/go-rest-api/db"
	"github.com/vinicius-benevides/go-rest-api/models"
)

func GetAllEvents() ([]models.Event, error) {
	query := "SELECT * FROM events"

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Could not get all events: %v", err)
	}
	defer rows.Close()

	var events []models.Event

	for rows.Next() {
		var event models.Event
		if err = rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserID,
		); err != nil {
			return nil, fmt.Errorf("Could not get all events: %v", err)
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*models.Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event models.Event
	if err := row.Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.DateTime,
		&event.UserID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("Could not get event: %v", err)
	}

	return &event, nil
}

func CreateEvent(event *models.Event) error {
	query := `
        INSERT INTO events(name, description, location, dateTime, user_id) 
        VALUES (?, ?, ?, ?, ?)
    `

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("Could not prepare statement for saving event: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)
	if err != nil {
		return fmt.Errorf("Could not execute statement for saving event: %v", err)
	}

	event.ID, err = result.LastInsertId()
	if err != nil {
		return fmt.Errorf("Could not get last inserted id for saved event: %v", err)
	}

	return nil
}

func UpdateEvent(event *models.Event) error {
	query := `
        UPDATE events
        SET name = ?, description = ?, location = ?, dateTime = ? 
        WHERE id = ?
    `

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("Could not prepare statement for updating event: %v", err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID); err != nil {
		return fmt.Errorf("Could not execute statement for updating event: %v", err)
	}

	return nil
}

func DeleteEvent(id int64) error {
	query := `
        DELETE FROM events
        WHERE id = ?
    `

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("Could not prepare statement for deleting event: %v", err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(id); err != nil {
		return fmt.Errorf("Could not execute statement for deleting event: %v", err)
	}

	return nil
}
