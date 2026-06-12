package config

import (
	"fmt"
	"os"
)

type Config struct {
	HTTPAddr    string
	DatabaseURL string
}

func Load() (Config, error) {
	databaseURL, err := requiredEnv("DATABASE_URL")
	if err != nil {
		return Config{}, err
	}

	return Config{
		HTTPAddr:    env("HTTP_ADDR", ":8080"),
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
