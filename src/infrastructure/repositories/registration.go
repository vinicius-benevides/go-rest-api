package repositories

import (
	"github.com/vinicius-benevides/go-rest-api/pkg/errs"
	"github.com/vinicius-benevides/go-rest-api/src/infrastructure/db"
)

func GetRegistrationsByEvent(eventID int64) ([]int64, error) {
	query := "SELECT user_id FROM registrations WHERE event_id = ?"

	rows, err := db.DB.Query(query, eventID)
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

func CreateRegistration(eventID, userID int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return errs.InternalError("Could not create registration", err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(eventID, userID); err != nil {
		return errs.InternalError("Could not create registration", err)
	}

	return nil
}

func CancelRegistration(eventID, userID int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return errs.InternalError("Could not cancel registration", err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(eventID, userID); err != nil {
		return errs.InternalError("Could not cancel registration", err)
	}

	return nil
}
