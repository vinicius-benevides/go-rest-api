package models

import (
	"fmt"

	"github.com/vinicius-benevides/go-rest-api/db"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
		INSERT INTO users(email, password) 
		VALUES (?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("Could not prepare statement for saving user: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return fmt.Errorf("Could not execute statement for saving user: %v", err)
	}

	u.ID, err = result.LastInsertId()
	if err != nil {
		return fmt.Errorf("Could not get last inserted id for saved user: %v", err)
	}

	return nil
}
