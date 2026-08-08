package opensearch

import "github.com/opensearch-project/opensearch-go/v4"

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
