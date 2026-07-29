package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		return nil, errors.New("[LoadConfig]: HTTP_PORT is not set")
	}

	return &Config{
		Port: port,
	}, nil
}
