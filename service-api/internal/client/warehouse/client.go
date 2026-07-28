package warehouse

import (
	"context"
	"net/http"
	"service-api/internal/dto/request"
)

type Client struct {
	baseURL string
	client  *http.Client
}

func New(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func (w *Client) CreateProduct(ctx context.Context, req request.CreateProductRequest) error {
	return nil
}
