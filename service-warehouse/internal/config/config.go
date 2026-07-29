package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	DB   DBConfig
}

type DBConfig struct {
	Driver   string
	Host     string
	PortNum  int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		return nil, errors.New("[LoadConfig]: HTTP_PORT is not set")
	}

	dbPort := 5432
	if rawPort := os.Getenv("DB_PORT"); rawPort != "" {
		if parsed, err := strconv.Atoi(rawPort); err == nil {
			dbPort = parsed
		}
	}

	return &Config{
		Port: port,
		DB: DBConfig{
			Driver:   getEnv("DB_DRIVER", "postgres"),
			Host:     getEnv("DB_HOST", "localhost"),
			PortNum:  dbPort,
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "warehouse"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
	}, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
