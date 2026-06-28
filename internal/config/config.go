package config

import (
	"fmt"
	"os"
)

const (
	envDatabaseURL = "DATABASE_URL"
	envHTTPAddr    = "HTTP_ADDR"

	defaultHTTPAddr = ":8080"
)

type Config struct {
	HTTPAddr    string
	DatabaseURL string
}

func Load() (Config, error) {
	databaseURL, err := requiredEnv(envDatabaseURL)
	if err != nil {
		return Config{}, err
	}

	return Config{
		HTTPAddr:    env(envHTTPAddr, defaultHTTPAddr),
		DatabaseURL: databaseURL,
	}, nil
}

func env(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return defaultValue
}

func requiredEnv(key string) (string, error) {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value, nil
	}

	return "", fmt.Errorf("environment variable %s is required but not set", key)
}
