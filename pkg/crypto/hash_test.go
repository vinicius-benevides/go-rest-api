package crypto

import "testing"

func TestHashPasswordAndCheck(t *testing.T) {
	hashed, err := HashPassword("super-secret")
	if err != nil {
		t.Fatalf("unexpected error hashing password: %v", err)
	}
	if hashed == "super-secret" {
		t.Fatalf("expected hashed password to differ from plain text")
	}

	if !CheckPasswordHash("super-secret", hashed) {
		t.Fatalf("expected password to validate")
	}

	if CheckPasswordHash("wrong", hashed) {
		t.Fatalf("expected validation to fail for wrong password")
	}
}
