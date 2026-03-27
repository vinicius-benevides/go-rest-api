package services

import (
	"github.com/vinicius-benevides/go-rest-api/pkg/errs"
	"github.com/vinicius-benevides/go-rest-api/src/models"
	"github.com/vinicius-benevides/go-rest-api/utils"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
}

type UserService interface {
	Signup(user models.User) (*models.User, error)
	Login(email, password string) (string, error)
}

type userService struct {
	repo         UserRepository
	tokenService TokenService
}

func NewUserService(repo UserRepository, tokenService TokenService) UserService {
	return &userService{repo: repo, tokenService: tokenService}
}

func (s *userService) Signup(user models.User) (*models.User, error) {
	if _, err := s.repo.GetUserByEmail(user.Email); err == nil {
		return nil, errs.ConflictError("Email already registered")
	} else if !errs.Is(err, errs.NotFound) {
		return nil, err
	}

	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, errs.InternalError("Could not process user data", err)
	}

	user.Password = hashed
	if err := s.repo.CreateUser(&user); err != nil {
		return nil, err
	}

	user.Password = ""
	return &user, nil
}

func (s *userService) Login(email, password string) (string, error) {
	foundUser, err := s.repo.GetUserByEmail(email)
	if err != nil {
		if errs.Is(err, errs.NotFound) {
			return "", errs.UnauthorizedError("Invalid credentials")
		}
		return "", err
	}

	if !utils.CheckPasswordHash(password, foundUser.Password) {
		return "", errs.UnauthorizedError("Invalid credentials")
	}

	token, err := s.tokenService.GenerateToken(foundUser.ID, foundUser.Email)
	if err != nil {
		return "", errs.InternalError("Could not generate authentication token", err)
	}

	return token, nil
}
