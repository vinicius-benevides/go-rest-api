package repositories

import (
	"database/sql"

	"github.com/vinicius-benevides/go-rest-api/pkg/errs"
)

type RegistrationRepository struct {
	db *sql.DB
}

func NewRegistrationRepository(db *sql.DB) *RegistrationRepository {
	return &RegistrationRepository{db: db}
}

func (r *RegistrationRepository) GetRegistrationsByEvent(eventID int64) ([]int64, error) {
	query := "SELECT user_id FROM registrations WHERE event_id = ?"

	rows, err := r.db.Query(query, eventID)
	if err != nil {
		return nil, errs.InternalError("Could not get registrations", err)
	}
	defer rows.Close()

	var users []int64
	for rows.Next() {
		var userID int64
		if err = rows.Scan(&userID); err != nil {
			return nil, errs.InternalError("Could not get registrations", err)
		}
		users = append(users, userID)
	}

	return users, nil
}

func (r *RegistrationRepository) CreateRegistration(eventID, userID int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return errs.InternalError("Could not create registration", err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(eventID, userID); err != nil {
		return errs.InternalError("Could not create registration", err)
	}

	return nil
}

func (r *RegistrationRepository) CancelRegistration(eventID, userID int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return errs.InternalError("Could not cancel registration", err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(eventID, userID); err != nil {
		return errs.InternalError("Could not cancel registration", err)
	}

	return nil
}
