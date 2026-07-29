package warehouse

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"service-api/internal/dto/request"
	"service-api/internal/dto/warehouse"
	apperrors "service-api/internal/errors"
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
		return fmt.Errorf("create request: %w", err)
	}
	resp, err := w.client.Do(data)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response body: %w", err)
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		var apiErr struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		}
		if err := json.Unmarshal(respBody, &apiErr); err == nil && apiErr.Message != "" {
			return apperrors.NewUnexpectedStatusError(resp.StatusCode, apiErr.Message)
		}
		return apperrors.NewUnexpectedStatusError(resp.StatusCode, string(respBody))
	}

	return nil
}
