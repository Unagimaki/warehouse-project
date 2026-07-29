package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"service-api/internal/dto/request"
	apperrors "service-api/internal/errors"
)

type ProductHandler struct {
	warehouse WarehouseClient
}

func New(warehouseClient WarehouseClient) *ProductHandler {
	return &ProductHandler{
		warehouse: warehouseClient,
	}
}

type WarehouseClient interface {
	CreateProduct(ctx context.Context, req request.CreateProductRequest) error
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req request.CreateProductRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if err := h.warehouse.CreateProduct(r.Context(), req); err != nil {
		var appErr *apperrors.AppError
		if errors.As(err, &appErr) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(appErr.StatusCode)
			_ = json.NewEncoder(w).Encode(map[string]string{"code": appErr.Code, "message": appErr.Message})
			return
		}

		log.Printf("create product failed: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{"code": "internal_error", "message": "internal server error"})
		return
	}
}
