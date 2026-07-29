package errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code       string
	Message    string
	StatusCode int
	Cause      error
}

func (e *AppError) Error() string {
	if e == nil {
		return ""
	}
	if e.Message != "" {
		return e.Message
	}
	return e.Code
}

func (e *AppError) Unwrap() error {
	return e.Cause
}

func NewBadRequestError(message string) *AppError {
	return &AppError{Code: "bad_request", Message: message, StatusCode: http.StatusBadRequest}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{Code: "not_found", Message: message, StatusCode: http.StatusNotFound}
}

func NewConflictError(message string) *AppError {
	return &AppError{Code: "conflict", Message: message, StatusCode: http.StatusConflict}
}

func NewInternalError(message string, cause error) *AppError {
	return &AppError{Code: "internal_error", Message: message, StatusCode: http.StatusInternalServerError, Cause: cause}
}

func NewUnexpectedStatusError(statusCode int, body string) error {
	return &AppError{Code: "warehouse_error", Message: body, StatusCode: statusCode}
}

func NewValidationError(message string) error {
	return fmt.Errorf("validation failed: %w", NewBadRequestError(message))
}
