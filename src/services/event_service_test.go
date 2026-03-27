package services

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/vinicius-benevides/go-rest-api/pkg/errs"
	"github.com/vinicius-benevides/go-rest-api/src/domain/models"
	"github.com/vinicius-benevides/go-rest-api/test/mocks"
)

func TestEventService(t *testing.T) {
	t.Run("ListEvents", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockEventRepository(ctrl)
		repo.EXPECT().GetAllEvents().Return([]models.Event{{ID: 1}}, nil)

		service := NewEventService(repo)
		events, err := service.ListEvents()
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if len(events) != 1 || events[0].ID != 1 {
			t.Fatalf("unexpected events %v", events)
		}
	})

	t.Run("CreateEventSetsOwner", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockEventRepository(ctrl)
		repo.EXPECT().CreateEvent(gomock.Any()).DoAndReturn(func(event *models.Event) error {
			if event.UserID != 10 {
				t.Fatalf("expected owner 10, got %d", event.UserID)
			}
			event.ID = 99
			return nil
		})

		service := NewEventService(repo)
		created, err := service.CreateEvent(10, models.Event{Name: "Launch"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if created.ID != 99 {
			t.Fatalf("expected id to be set, got %d", created.ID)
		}
	})

	t.Run("UpdateEventUnauthorized", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockEventRepository(ctrl)
		repo.EXPECT().GetEventByID(int64(1)).Return(&models.Event{ID: 1, UserID: 5}, nil)

		service := NewEventService(repo)
		_, err := service.UpdateEvent(10, 1, models.Event{Name: "Update"})
		if err == nil || !errs.Is(err, errs.Unauthorized) {
			t.Fatalf("expected unauthorized error, got %v", err)
		}
	})

	t.Run("DeleteEventPropagatesRepoError", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockEventRepository(ctrl)
		repo.EXPECT().GetEventByID(int64(1)).Return(&models.Event{ID: 1, UserID: 2}, nil)
		repo.EXPECT().DeleteEvent(int64(1)).Return(errors.New("boom"))

		service := NewEventService(repo)
		if err := service.DeleteEvent(2, 1); err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}
