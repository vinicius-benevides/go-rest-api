package models

import (
	"database/sql"
	"fmt"

	"github.com/vinicius-benevides/go-rest-api/db"
	"github.com/vinicius-benevides/go-rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `json:",omitempty" binding:"required"`
}

func GetUserByEmail(email string) (*User, error) {
	query := "SELECT * FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, email)

	var user User
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

func (u *User) Login() error {
	foundUser, err := GetUserByEmail(u.Email)
	if err != nil || foundUser == nil {
		return fmt.Errorf("Invalid credentials")
	}

	validPassword := utils.CheckPasswordHash(u.Password, foundUser.Password)

	if !validPassword {
		return fmt.Errorf("Invalid credentials")
	}

	return nil
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

	u.Password, err = utils.HashPassword(u.Password)
	if err != nil {
		return fmt.Errorf("Could not hash user password: %v", err)
	}

	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return fmt.Errorf("Could not execute statement for saving user: %v", err)
	}

	u.ID, err = result.LastInsertId()
	if err != nil {
		return fmt.Errorf("Could not get last inserted id for saved user: %v", err)
	}

	u.Password = ""

	return nil
}
