package opensearch

import (
	"bytes"
	"context"
	"encoding/json"
	"service-analytics/internal/domain"
	"strconv"

	"github.com/opensearch-project/opensearch-go/v4"
	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
)

type Product struct {
	ID          int    `json:"id"`
	ItemNumber  string `json:"item_number"`
	Type        string `json:"type"`
	Brand       string `json:"brand"`
	Gender      string `json:"gender"`
	VolumeML    int    `json:"volume_ml"`
	Description string `json:"description"`
	Barcode     string `json:"barcode"`
	CreatedAt   int64  `json:"created_at"`
}

type Client struct {
	client *opensearch.Client
}

func NewClient() (*Client, error) {
	client, err := opensearch.NewClient(opensearch.Config{
		Addresses: []string{
			"http://opensearch:9200",
		},
	})
	if err != nil {
		return nil, err
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) IndexProduct(ctx context.Context, product domain.Product) error {
	body, err := json.Marshal(product)
	if err != nil {
		return err
	}

	req := opensearchapi.IndexReq{
		Index:      "products",
		DocumentID: strconv.Itoa(product.ID),
		Body:       bytes.NewReader(body),
	}

	_, err = c.client.Do(ctx, "PUT", req, nil)
	if err != nil {
		return err
	}

	return nil
}
