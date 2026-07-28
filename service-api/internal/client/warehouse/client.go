package warehouse

import (
	"context"
	"net/http"
	"service-api/internal/dto/request"
)

type WarehouseClient struct {
	baseURL string
	client  *http.Client
}

func New(baseURL string) *WarehouseClient {
	return &WarehouseClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func (w *WarehouseClient) CreateProduct(ctx context.Context, req request.CreateProductRequest) error {
	return nil
}
