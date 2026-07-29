package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
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

func WriteJSONError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	var appErr *AppError
	if errors.As(err, &appErr) {
		log.Printf("[warehouse] code=%s status=%d message=%s cause=%v", appErr.Code, appErr.StatusCode, appErr.Message, appErr.Cause)
		w.WriteHeader(appErr.StatusCode)
		_ = json.NewEncoder(w).Encode(map[string]string{"code": appErr.Code, "message": appErr.Message})
		return
	}

	log.Printf("[warehouse] unexpected error: %v", err)
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(map[string]string{"code": "internal_error", "message": "internal server error"})
}

func Wrap(message string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", message, err)
}
