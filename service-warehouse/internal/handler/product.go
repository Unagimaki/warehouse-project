package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"service-werehouse/internal/dto/request"
	apperrors "service-werehouse/internal/errors"
)

type ProductService interface {
	CreateProduct(ctx context.Context, req request.CreateProductRequest) error
}

type ProductHandler struct {
	service ProductService
}

func NewProductHandler(service ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req request.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		apperrors.WriteJSONError(w, apperrors.NewBadRequestError("invalid request body"))
		return
	}

	if err := h.service.CreateProduct(r.Context(), req); err != nil {
		apperrors.WriteJSONError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(`{"message":"created"}`))
}
