package services

import (
	"github.com/vinicius-benevides/go-rest-api/pkg/errs"
	"github.com/vinicius-benevides/go-rest-api/src/domain/models"
	"github.com/vinicius-benevides/go-rest-api/src/domain/ports"
)

type EventService interface {
	ListEvents() ([]models.Event, error)
	GetEvent(id int64) (*models.Event, error)
	CreateEvent(userID int64, event models.Event) (*models.Event, error)
	UpdateEvent(userID, eventID int64, updated models.Event) (*models.Event, error)
	DeleteEvent(userID, eventID int64) error
}

type eventService struct {
	repo ports.EventRepository
}

func NewEventService(repo ports.EventRepository) EventService {
	return &eventService{repo: repo}
}

func (s *eventService) ListEvents() ([]models.Event, error) {
	return s.repo.GetAllEvents()
}

func (s *eventService) GetEvent(id int64) (*models.Event, error) {
	return s.repo.GetEventByID(id)
}

func (s *eventService) CreateEvent(userID int64, event models.Event) (*models.Event, error) {
	event.UserID = userID
	if err := s.repo.CreateEvent(&event); err != nil {
		return nil, err
	}
	return &event, nil
}

func (s *eventService) UpdateEvent(userID, eventID int64, updated models.Event) (*models.Event, error) {
	existing, err := s.repo.GetEventByID(eventID)
	if err != nil {
		return nil, err
	}

	if existing.UserID != userID {
		return nil, errs.UnauthorizedError("Not authorized to update event")
	}

	updated.ID = eventID
	updated.UserID = userID

	if err := s.repo.UpdateEvent(&updated); err != nil {
		return nil, err
	}

	return &updated, nil
}

func (s *eventService) DeleteEvent(userID, eventID int64) error {
	existing, err := s.repo.GetEventByID(eventID)
	if err != nil {
		return err
	}

	if existing.UserID != userID {
		return errs.UnauthorizedError("Not authorized to delete event")
	}

	return s.repo.DeleteEvent(eventID)
}
