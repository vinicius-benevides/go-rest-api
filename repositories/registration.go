package repositories

import (
	"fmt"

	"github.com/vinicius-benevides/go-rest-api/db"
)

func GetRegistrationsByEvent(eventID int64) ([]int64, error) {
	query := "SELECT user_id FROM registrations WHERE event_id = ?"

	rows, err := db.DB.Query(query, eventID)
	if err != nil {
		return nil, fmt.Errorf("Could not get registrations: %v", err)
	}
	defer rows.Close()

	var users []int64
	for rows.Next() {
		var userID int64
		if err = rows.Scan(&userID); err != nil {
			return nil, fmt.Errorf("Could not get registrations: %v", err)
		}
		users = append(users, userID)
	}

	return users, nil
}

func CreateRegistration(eventID, userID int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("Could not prepare statement for creating registration for event: %v", err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(eventID, userID); err != nil {
		return fmt.Errorf("Could not execute statement for creating registration for event: %v", err)
	}

	return nil
}

func CancelRegistration(eventID, userID int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("Could not prepare statement for canceling registration for event: %v", err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(eventID, userID); err != nil {
		return fmt.Errorf("Could not execute statement for canceling registration for event: %v", err)
	}

	return nil
}
