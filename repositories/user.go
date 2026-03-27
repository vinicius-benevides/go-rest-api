package repositories

import (
	"database/sql"
	"fmt"

	"github.com/vinicius-benevides/go-rest-api/db"
	"github.com/vinicius-benevides/go-rest-api/models"
	"github.com/vinicius-benevides/go-rest-api/utils"
)

func GetUserByEmail(email string) (*models.User, error) {
	query := "SELECT * FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, email)

	var user models.User
	if err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("Could not get user: %v", err)
	}

	return &user, nil
}

func CreateUser(user *models.User) error {
	query := `
        INSERT INTO users(email, password) 
        VALUES (?, ?)
    `

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("Could not prepare statement for saving user: %v", err)
	}
	defer stmt.Close()

	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("Could not hash user password: %v", err)
	}

	result, err := stmt.Exec(user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("Could not execute statement for saving user: %v", err)
	}

	user.ID, err = result.LastInsertId()
	if err != nil {
		return fmt.Errorf("Could not get last inserted id for saved user: %v", err)
	}

	user.Password = ""

	return nil
}
