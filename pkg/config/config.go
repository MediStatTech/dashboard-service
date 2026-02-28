package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	// Server configuration
	ServerHost string `env:"SERVER_HOST"`
	ServerPort string `env:"SERVER_PORT"`

	// TLS configuration
	TLSCertFilePath string `env:"TLS_CERT_FILE_PATH"`
	TLSKeyFilePath  string `env:"TLS_KEY_FILE_PATH"`

	// JWT configuration
	JWTSecretKey string        `env:"JWT_SECRET_KEY"`
	JWTDuration  time.Duration `env:"JWT_DURATION"`

	// Environment
	Environment string `env:"ENVIRONMENT"`

	// Logging
	LogLevel string `env:"LOG_LEVEL"`
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
	cfg.JWTSecretKey = getEnv("JWT_SECRET_KEY", "")
	cfg.JWTDuration = getEnvAsDuration("JWT_DURATION", 24*time.Hour)
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

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := time.ParseDuration(valueStr)
	if err != nil {
		return defaultValue
	}

	return value
}
