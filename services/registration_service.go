package services

import (
	"slices"

	"github.com/vinicius-benevides/go-rest-api/pkg/errs"
	"github.com/vinicius-benevides/go-rest-api/repositories"
)

type RegistrationService struct{}

func NewRegistrationService() *RegistrationService {
	return &RegistrationService{}
}

func (s *RegistrationService) RegisterUser(eventID, userID int64) error {
	if _, err := repositories.GetEventByID(eventID); err != nil {
		return err
	}

	users, err := repositories.GetRegistrationsByEvent(eventID)
	if err != nil {
		return err
	}

	if slices.Contains(users, userID) {
		return errs.ConflictError("User already registered for this event")
	}

	return repositories.CreateRegistration(eventID, userID)
}

func (s *RegistrationService) CancelRegistration(eventID, userID int64) error {
	if _, err := repositories.GetEventByID(eventID); err != nil {
		return err
	}

	users, err := repositories.GetRegistrationsByEvent(eventID)
	if err != nil {
		return err
	}

	if !slices.Contains(users, userID) {
		return errs.NotFoundError("User is not registered for this event")
	}

	return repositories.CancelRegistration(eventID, userID)
}
