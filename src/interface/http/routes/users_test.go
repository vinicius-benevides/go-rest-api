package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/vinicius-benevides/go-rest-api/pkg/errs"
	"github.com/vinicius-benevides/go-rest-api/src/domain/models"
	"github.com/vinicius-benevides/go-rest-api/test/mocks"
)

func TestUserRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("SignupInvalidPayload", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		handler := &Handler{userService: mocks.NewMockUserService(ctrl)}
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBufferString("invalid"))

		handler.signup(ctx)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", w.Code)
		}
	})

	t.Run("SignupSuccess", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := mocks.NewMockUserService(ctrl)
		handler := &Handler{userService: service}

		service.EXPECT().Signup(models.User{Email: "test@test", Password: "pw"}).Return(&models.User{ID: 1}, nil)

		body := bytes.NewBufferString(`{"email":"test@test","password":"pw"}`)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, "/signup", body)
		ctx.Request.Header.Set("Content-Type", "application/json")

		handler.signup(ctx)

		if w.Code != http.StatusCreated {
			t.Fatalf("expected 201, got %d", w.Code)
		}

		var resp map[string]any
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			t.Fatalf("failed to parse response: %v", err)
		}
		if resp["message"] != "User created" {
			t.Fatalf("unexpected message %v", resp["message"])
		}
	})

	t.Run("LoginUnauthorized", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := mocks.NewMockUserService(ctrl)
		handler := &Handler{userService: service}

		service.EXPECT().Login("test@test", "pw").Return("", errs.UnauthorizedError("Invalid credentials"))

		body := bytes.NewBufferString(`{"email":"test@test","password":"pw"}`)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, "/login", body)
		ctx.Request.Header.Set("Content-Type", "application/json")

		handler.login(ctx)

		if w.Code != http.StatusUnauthorized {
			t.Fatalf("expected 401, got %d", w.Code)
		}
	})
}
