package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// Server configuration
	ServerHost string
	ServerPort string

	// TLS configuration
	TLSCertFilePath string
	TLSKeyFilePath  string

	// Environment
	Environment string

	// Logging
	LogLevel string
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("error loading .env file: %w", err)
		}
	}

	cfg := &Config{}

	cfg.ServerHost = getEnv("SERVER_HOST", "0.0.0.0")
	cfg.ServerPort = getEnv("SERVER_PORT", "8080")
	cfg.TLSCertFilePath = getEnv("TLS_CERT_FILE_PATH", "")
	cfg.TLSKeyFilePath = getEnv("TLS_KEY_FILE_PATH", "")
	cfg.Environment = getEnv("ENVIRONMENT", "development")
	cfg.LogLevel = getEnv("LOG_LEVEL", "info")

	return cfg, nil
}

func getEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
