package services

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/vinicius-benevides/go-rest-api/pkg/crypto"
	"github.com/vinicius-benevides/go-rest-api/pkg/errs"
	"github.com/vinicius-benevides/go-rest-api/src/domain/models"
	"github.com/vinicius-benevides/go-rest-api/test/mocks"
)

func TestUserService(t *testing.T) {
	t.Run("SignupEmailExists", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockUserRepository(ctrl)
		token := mocks.NewMockTokenService(ctrl)

		repo.EXPECT().GetUserByEmail("a@test").Return(&models.User{ID: 1}, nil)

		service := NewUserService(repo, token)
		if _, err := service.Signup(models.User{Email: "a@test", Password: "pw"}); err == nil || !errs.Is(err, errs.Conflict) {
			t.Fatalf("expected conflict, got %v", err)
		}
	})

	t.Run("SignupSuccess", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockUserRepository(ctrl)
		token := mocks.NewMockTokenService(ctrl)

		repo.EXPECT().GetUserByEmail("new@test").Return(nil, errs.NotFoundError("not found"))
		repo.EXPECT().CreateUser(gomock.Any()).DoAndReturn(func(user *models.User) error {
			if user.Password == "plain" {
				t.Fatalf("expected password to be hashed")
			}
			user.ID = 10
			return nil
		})

		service := NewUserService(repo, token)
		created, err := service.Signup(models.User{Email: "new@test", Password: "plain"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if created.ID != 10 || created.Password != "" {
			t.Fatalf("expected response sanitized, got %+v", created)
		}
	})

	t.Run("LoginInvalidPassword", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockUserRepository(ctrl)
		token := mocks.NewMockTokenService(ctrl)

		hashed, _ := crypto.HashPassword("correct")
		repo.EXPECT().GetUserByEmail("test@test").Return(&models.User{ID: 1, Email: "test@test", Password: hashed}, nil)

		service := NewUserService(repo, token)
		if _, err := service.Login("test@test", "wrong"); err == nil || !errs.Is(err, errs.Unauthorized) {
			t.Fatalf("expected unauthorized, got %v", err)
		}
	})

	t.Run("LoginSuccess", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockUserRepository(ctrl)
		token := mocks.NewMockTokenService(ctrl)

		hashed, _ := crypto.HashPassword("correct")
		repo.EXPECT().GetUserByEmail("test@test").Return(&models.User{ID: 42, Email: "test@test", Password: hashed}, nil)
		token.EXPECT().GenerateToken(int64(42), "test@test").Return("signed-token", nil)

		service := NewUserService(repo, token)
		tok, err := service.Login("test@test", "correct")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if tok != "signed-token" {
			t.Fatalf("unexpected token %s", tok)
		}
	})
}
