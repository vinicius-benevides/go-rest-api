package services

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/vinicius-benevides/go-rest-api/pkg/errs"
	"github.com/vinicius-benevides/go-rest-api/src/domain/models"
	"github.com/vinicius-benevides/go-rest-api/test/mocks"
)

func TestRegistrationService(t *testing.T) {
	t.Run("RegisterConflict", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		eventRepo := mocks.NewMockEventRepository(ctrl)
		regRepo := mocks.NewMockRegistrationRepository(ctrl)

		eventRepo.EXPECT().GetEventByID(int64(1)).Return(&models.Event{ID: 1}, nil)
		regRepo.EXPECT().GetRegistrationsByEvent(int64(1)).Return([]int64{2}, nil)

		service := NewRegistrationService(eventRepo, regRepo)
		if err := service.RegisterUser(1, 2); err == nil || !errs.Is(err, errs.Conflict) {
			t.Fatalf("expected conflict, got %v", err)
		}
	})

	t.Run("CancelNotFound", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		eventRepo := mocks.NewMockEventRepository(ctrl)
		regRepo := mocks.NewMockRegistrationRepository(ctrl)

		eventRepo.EXPECT().GetEventByID(int64(1)).Return(&models.Event{ID: 1}, nil)
		regRepo.EXPECT().GetRegistrationsByEvent(int64(1)).Return([]int64{}, nil)

		service := NewRegistrationService(eventRepo, regRepo)
		if err := service.CancelRegistration(1, 2); err == nil || !errs.Is(err, errs.NotFound) {
			t.Fatalf("expected not found, got %v", err)
		}
	})

	t.Run("RegisterSuccess", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		eventRepo := mocks.NewMockEventRepository(ctrl)
		regRepo := mocks.NewMockRegistrationRepository(ctrl)

		eventRepo.EXPECT().GetEventByID(int64(2)).Return(&models.Event{ID: 2}, nil)
		regRepo.EXPECT().GetRegistrationsByEvent(int64(2)).Return([]int64{}, nil)
		regRepo.EXPECT().CreateRegistration(int64(2), int64(3)).Return(nil)

		service := NewRegistrationService(eventRepo, regRepo)
		if err := service.RegisterUser(2, 3); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
