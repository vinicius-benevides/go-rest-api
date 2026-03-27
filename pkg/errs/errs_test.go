package errs

import (
	"errors"
	"testing"
)

func TestAppErrorBasics(t *testing.T) {
	inner := errors.New("boom")
	err := New(Internal, "something broke", inner)

	if err.Code() != Internal {
		t.Fatalf("expected code %s, got %s", Internal, err.Code())
	}

	if err.Error() != "something broke" {
		t.Fatalf("unexpected message %q", err.Error())
	}

	if !errors.Is(err, inner) {
		t.Fatalf("expected errors.Is to unwrap inner")
	}

	if !Is(err, Internal) {
		t.Fatalf("expected Is to match internal code")
	}

	if Is(err, NotFound) {
		t.Fatalf("did not expect not found match")
	}
}

func TestHelperConstructors(t *testing.T) {
	tests := []struct {
		name string
		err  *AppError
		code Code
	}{
		{"notfound", NotFoundError("missing"), NotFound},
		{"conflict", ConflictError("dup"), Conflict},
		{"unauthorized", UnauthorizedError("nope"), Unauthorized},
		{"badrequest", BadRequestError("bad"), BadRequest},
	}

	for _, tt := range tests {
		if tt.err.Code() != tt.code {
			t.Fatalf("%s expected code %s got %s", tt.name, tt.code, tt.err.Code())
		}
		if tt.err.Error() == "" {
			t.Fatalf("%s expected message", tt.name)
		}
	}

	wrapped := InternalError("err", errors.New("root"))
	if wrapped.Code() != Internal {
		t.Fatalf("internal helper returned wrong code")
	}
}
