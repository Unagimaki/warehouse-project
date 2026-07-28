package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"service-api/internal/dto/request"
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
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
}
