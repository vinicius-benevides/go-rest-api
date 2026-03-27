package repositories

import (
	"database/sql"

	"github.com/vinicius-benevides/go-rest-api/pkg/errs"
	"github.com/vinicius-benevides/go-rest-api/src/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	query := "SELECT * FROM users WHERE email = ?"
	row := r.db.QueryRow(query, email)

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

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `
        INSERT INTO users(email, password) 
        VALUES (?, ?)
    `

	stmt, err := r.db.Prepare(query)
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
