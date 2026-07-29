package warehouse

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"service-api/internal/dto/request"
	"service-api/internal/dto/warehouse"
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
	var warehouseReq = warehouse.CreateProductRequest{
		ItemNumber: req.Barcode,
		Type:       req.Category,
		Brand:      req.Brand,
		Meta: warehouse.ProductMeta{
			Gender:      req.Gender,
			VolumeML:    req.VolumeML,
			Description: req.Description,
			Barcode:     req.Barcode,
		},
	}

	body, err := json.Marshal(warehouseReq)
	if err != nil {
		return err
	}
	data, err := http.NewRequestWithContext(ctx, http.MethodPost, w.baseURL+"/products", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	resp, err := w.client.Do(data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return errors.New("warehouse returned unexpected status")
	}
	return nil
}
