package services

import (
	"github.com/vinicius-benevides/go-rest-api/pkg/errs"
	"github.com/vinicius-benevides/go-rest-api/src/infrastructure/repositories"
	"github.com/vinicius-benevides/go-rest-api/src/models"
)

type EventService struct{}

func NewEventService() *EventService {
	return &EventService{}
}

func (s *EventService) ListEvents() ([]models.Event, error) {
	return repositories.GetAllEvents()
}

func (s *EventService) GetEvent(id int64) (*models.Event, error) {
	return repositories.GetEventByID(id)
}

func (s *EventService) CreateEvent(userID int64, event models.Event) (*models.Event, error) {
	event.UserID = userID
	if err := repositories.CreateEvent(&event); err != nil {
		return nil, err
	}
	return &event, nil
}

func (s *EventService) UpdateEvent(userID, eventID int64, updated models.Event) (*models.Event, error) {
	existing, err := repositories.GetEventByID(eventID)
	if err != nil {
		return nil, err
	}

	if existing.UserID != userID {
		return nil, errs.UnauthorizedError("Not authorized to update event")
	}

	updated.ID = eventID
	updated.UserID = userID

	if err := repositories.UpdateEvent(&updated); err != nil {
		return nil, err
	}

	return &updated, nil
}

func (s *EventService) DeleteEvent(userID, eventID int64) error {
	existing, err := repositories.GetEventByID(eventID)
	if err != nil {
		return err
	}

	if existing.UserID != userID {
		return errs.UnauthorizedError("Not authorized to delete event")
	}

	return repositories.DeleteEvent(eventID)
}
