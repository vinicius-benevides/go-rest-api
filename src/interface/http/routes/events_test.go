package routes

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/vinicius-benevides/go-rest-api/pkg/errs"
	"github.com/vinicius-benevides/go-rest-api/src/domain/models"
	"github.com/vinicius-benevides/go-rest-api/test/mocks"
)

func TestEventRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("GetEventsSuccess", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := mocks.NewMockEventService(ctrl)
		handler := &Handler{eventService: service}

		service.EXPECT().ListEvents().Return([]models.Event{{ID: 1}}, nil)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodGet, "/events", nil)

		handler.getEvents(ctx)

		if w.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", w.Code)
		}
	})

	t.Run("GetEventNotFound", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := mocks.NewMockEventService(ctrl)
		handler := &Handler{eventService: service}

		service.EXPECT().GetEvent(int64(1)).Return(nil, errs.NotFoundError("Event not found"))

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodGet, "/events/1", nil)
		ctx.Params = gin.Params{{Key: "id", Value: "1"}}

		handler.getEventByID(ctx)

		if w.Code != http.StatusNotFound {
			t.Fatalf("expected 404, got %d", w.Code)
		}
	})

	t.Run("CreateEventSuccess", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := mocks.NewMockEventService(ctrl)
		handler := &Handler{eventService: service}

		service.EXPECT().CreateEvent(int64(5), gomock.AssignableToTypeOf(models.Event{})).Return(&models.Event{ID: 99}, nil)

		body := bytes.NewBufferString(`{"name":"Launch","description":"desc","location":"loc","dateTime":"2026-03-26T00:00:00Z"}`)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, "/events", body)
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Set("userId", int64(5))

		handler.createEvent(ctx)

		if w.Code != http.StatusCreated {
			t.Fatalf("expected 201, got %d", w.Code)
		}
	})

	t.Run("UpdateEventUnauthorized", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := mocks.NewMockEventService(ctrl)
		handler := &Handler{eventService: service}

		service.EXPECT().UpdateEvent(int64(1), int64(2), gomock.AssignableToTypeOf(models.Event{})).Return(nil, errs.UnauthorizedError("Not authorized"))

		body := bytes.NewBufferString(`{"name":"New","description":"desc","location":"loc","dateTime":"2026-03-26T00:00:00Z"}`)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPut, "/events/2", body)
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Params = gin.Params{{Key: "id", Value: "2"}}
		ctx.Set("userId", int64(1))

		handler.updateEvent(ctx)

		if w.Code != http.StatusUnauthorized {
			t.Fatalf("expected 401, got %d", w.Code)
		}
	})

	t.Run("DeleteEventSuccess", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := mocks.NewMockEventService(ctrl)
		handler := &Handler{eventService: service}

		service.EXPECT().DeleteEvent(int64(1), int64(3)).Return(nil)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodDelete, "/events/3", nil)
		ctx.Params = gin.Params{{Key: "id", Value: "3"}}
		ctx.Set("userId", int64(1))

		handler.deleteEvent(ctx)

		if w.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", w.Code)
		}
	})
}
