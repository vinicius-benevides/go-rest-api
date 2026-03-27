package services

import (
	"testing"
	"time"
)

func TestTokenServiceGenerateAndVerify(t *testing.T) {
	service := NewTokenService("secret", time.Minute)
	token, err := service.GenerateToken(7, "user@test")
	if err != nil {
		t.Fatalf("unexpected error generating token: %v", err)
	}

	userID, err := service.VerifyToken(token)
	if err != nil {
		t.Fatalf("unexpected error verifying token: %v", err)
	}
	if userID != 7 {
		t.Fatalf("expected user id 7, got %d", userID)
	}
}
