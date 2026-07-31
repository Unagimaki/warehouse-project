package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	WarehouseURL string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	port := os.Getenv("HTTP_PORT")
	warehouseURL := os.Getenv("WAREHOUSE_URL")
	if port == "" {
		return nil, errors.New("[LoadConfig]: HTTP_PORT is not set")
	}
	if warehouseURL == "" {
		return nil, errors.New("[LoadConfig]: warehouseURL is not set")
	}

	return &Config{
		Port:         port,
		WarehouseURL: warehouseURL,
	}, nil
}
