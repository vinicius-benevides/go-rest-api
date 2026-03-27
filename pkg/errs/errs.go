package errs

import "errors"

// Code represents a domain level error category shared across layers.
type Code string

const (
	Internal     Code = "INTERNAL"
	NotFound     Code = "NOT_FOUND"
	Conflict     Code = "CONFLICT"
	Unauthorized Code = "UNAUTHORIZED"
	BadRequest   Code = "BAD_REQUEST"
)

// AppError captures the error category, a user friendly message, and an optional inner error.
type AppError struct {
	code    Code
	message string
	err     error
}

// New builds a new application error with the given metadata.
func New(code Code, message string, err error) *AppError {
	return &AppError{code: code, message: message, err: err}
}

// Error implements the error interface.
func (e *AppError) Error() string {
	return e.message
}

// Unwrap exposes the underlying error for logging or debugging when needed.
func (e *AppError) Unwrap() error {
	return e.err
}

// Code returns the error category.
func (e *AppError) Code() Code {
	return e.code
}

// Is compares an error against a specific AppError code.
func Is(err error, code Code) bool {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.code == code
	}
	return false
}

// Helpers for common error categories.
func NotFoundError(message string) *AppError {
	return New(NotFound, message, nil)
}

func ConflictError(message string) *AppError {
	return New(Conflict, message, nil)
}

func UnauthorizedError(message string) *AppError {
	return New(Unauthorized, message, nil)
}

func BadRequestError(message string) *AppError {
	return New(BadRequest, message, nil)
}

func InternalError(message string, err error) *AppError {
	return New(Internal, message, err)
}
