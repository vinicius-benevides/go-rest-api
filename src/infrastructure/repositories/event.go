package repositories

import (
	"database/sql"

	"github.com/vinicius-benevides/go-rest-api/pkg/errs"
	"github.com/vinicius-benevides/go-rest-api/src/models"
)

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) GetAllEvents() ([]models.Event, error) {
	query := "SELECT * FROM events"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, errs.InternalError("Could not get events", err)
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
			return nil, errs.InternalError("Could not get events", err)
		}

		events = append(events, event)
	}

	return events, nil
}

func (r *EventRepository) GetEventByID(id int64) (*models.Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := r.db.QueryRow(query, id)

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
			return nil, errs.NotFoundError("Event not found")
		}
		return nil, errs.InternalError("Could not get event", err)
	}

	return &event, nil
}

func (r *EventRepository) CreateEvent(event *models.Event) error {
	query := `
        INSERT INTO events(name, description, location, dateTime, user_id) 
        VALUES (?, ?, ?, ?, ?)
    `

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return errs.InternalError("Could not create event", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)
	if err != nil {
		return errs.InternalError("Could not create event", err)
	}

	event.ID, err = result.LastInsertId()
	if err != nil {
		return errs.InternalError("Could not create event", err)
	}

	return nil
}

func (r *EventRepository) UpdateEvent(event *models.Event) error {
	query := `
        UPDATE events
        SET name = ?, description = ?, location = ?, dateTime = ? 
        WHERE id = ?
    `

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return errs.InternalError("Could not update event", err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID); err != nil {
		return errs.InternalError("Could not update event", err)
	}

	return nil
}

func (r *EventRepository) DeleteEvent(id int64) error {
	query := `
        DELETE FROM events
        WHERE id = ?
    `

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return errs.InternalError("Could not delete event", err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(id); err != nil {
		return errs.InternalError("Could not delete event", err)
	}

	return nil
}
