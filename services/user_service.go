package services

import (
	"github.com/vinicius-benevides/go-rest-api/models"
	"github.com/vinicius-benevides/go-rest-api/pkg/errs"
	"github.com/vinicius-benevides/go-rest-api/repositories"
	"github.com/vinicius-benevides/go-rest-api/utils"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Signup(user models.User) (*models.User, error) {
	if _, err := repositories.GetUserByEmail(user.Email); err == nil {
		return nil, errs.ConflictError("Email already registered")
	} else if !errs.Is(err, errs.NotFound) {
		return nil, err
	}

	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, errs.InternalError("Could not process user data", err)
	}

	user.Password = hashed
	if err := repositories.CreateUser(&user); err != nil {
		return nil, err
	}

	user.Password = ""
	return &user, nil
}

func (s *UserService) Login(email, password string) (string, error) {
	foundUser, err := repositories.GetUserByEmail(email)
	if err != nil {
		if errs.Is(err, errs.NotFound) {
			return "", errs.UnauthorizedError("Invalid credentials")
		}
		return "", err
	}

	if !utils.CheckPasswordHash(password, foundUser.Password) {
		return "", errs.UnauthorizedError("Invalid credentials")
	}

	token, err := utils.GenerateToken(foundUser.ID, foundUser.Email)
	if err != nil {
		return "", errs.InternalError("Could not generate authentication token", err)
	}

	return token, nil
}
