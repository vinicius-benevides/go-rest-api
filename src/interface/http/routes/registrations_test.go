package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/vinicius-benevides/go-rest-api/pkg/errs"
	"github.com/vinicius-benevides/go-rest-api/test/mocks"
)

func TestRegistrationRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("CreateRegistrationSuccess", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := mocks.NewMockRegistrationService(ctrl)
		handler := &Handler{registrationService: service}

		service.EXPECT().RegisterUser(int64(4), int64(2)).Return(nil)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, "/events/4/register", nil)
		ctx.Params = gin.Params{{Key: "id", Value: "4"}}
		ctx.Set("userId", int64(2))

		handler.createRegistration(ctx)

		if w.Code != http.StatusCreated {
			t.Fatalf("expected 201, got %d", w.Code)
		}
	})

	t.Run("CreateRegistrationBadEventID", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		handler := &Handler{registrationService: mocks.NewMockRegistrationService(ctrl)}

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, "/events/x/register", nil)
		ctx.Params = gin.Params{{Key: "id", Value: "x"}}
		ctx.Set("userId", int64(1))

		handler.createRegistration(ctx)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", w.Code)
		}
	})

	t.Run("CancelRegistrationError", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := mocks.NewMockRegistrationService(ctrl)
		handler := &Handler{registrationService: service}

		service.EXPECT().CancelRegistration(int64(5), int64(9)).Return(errs.NotFoundError("not registered"))

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodDelete, "/events/5/register", nil)
		ctx.Params = gin.Params{{Key: "id", Value: "5"}}
		ctx.Set("userId", int64(9))

		handler.cancelRegistration(ctx)

		if w.Code != http.StatusNotFound {
			t.Fatalf("expected 404, got %d", w.Code)
		}
	})

	t.Run("CancelRegistrationSuccess", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := mocks.NewMockRegistrationService(ctrl)
		handler := &Handler{registrationService: service}

		service.EXPECT().CancelRegistration(int64(7), int64(3)).Return(nil)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodDelete, "/events/7/register", nil)
		ctx.Params = gin.Params{{Key: "id", Value: "7"}}
		ctx.Set("userId", int64(3))

		handler.cancelRegistration(ctx)

		if w.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", w.Code)
		}
	})
}
