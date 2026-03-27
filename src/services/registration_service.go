package services

import (
	"slices"

	"github.com/vinicius-benevides/go-rest-api/pkg/errs"
	"github.com/vinicius-benevides/go-rest-api/src/domain/ports"
)

type RegistrationService interface {
	RegisterUser(eventID, userID int64) error
	CancelRegistration(eventID, userID int64) error
}

type registrationService struct {
	eventRepo        ports.EventRepository
	registrationRepo ports.RegistrationRepository
}

func NewRegistrationService(eventRepo ports.EventRepository, registrationRepo ports.RegistrationRepository) RegistrationService {
	return &registrationService{
		eventRepo:        eventRepo,
		registrationRepo: registrationRepo,
	}
}

func (s *registrationService) RegisterUser(eventID, userID int64) error {
	if _, err := s.eventRepo.GetEventByID(eventID); err != nil {
		return err
	}

	users, err := s.registrationRepo.GetRegistrationsByEvent(eventID)
	if err != nil {
		return err
	}

	if slices.Contains(users, userID) {
		return errs.ConflictError("User already registered for this event")
	}

	return s.registrationRepo.CreateRegistration(eventID, userID)
}

func (s *registrationService) CancelRegistration(eventID, userID int64) error {
	if _, err := s.eventRepo.GetEventByID(eventID); err != nil {
		return err
	}

	users, err := s.registrationRepo.GetRegistrationsByEvent(eventID)
	if err != nil {
		return err
	}

	if !slices.Contains(users, userID) {
		return errs.NotFoundError("User is not registered for this event")
	}

	return s.registrationRepo.CancelRegistration(eventID, userID)
}
