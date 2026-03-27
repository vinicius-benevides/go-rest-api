package repositories

import (
	"database/sql"

	"github.com/vinicius-benevides/go-rest-api/db"
	"github.com/vinicius-benevides/go-rest-api/models"
	"github.com/vinicius-benevides/go-rest-api/pkg/errs"
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
			return nil, errs.NotFoundError("User not found")
		}
		return nil, errs.InternalError("Could not get user", err)
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
		return errs.InternalError("Could not create user", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Email, user.Password)
	if err != nil {
		return errs.InternalError("Could not create user", err)
	}

	user.ID, err = result.LastInsertId()
	if err != nil {
		return errs.InternalError("Could not create user", err)
	}

	return nil
}
