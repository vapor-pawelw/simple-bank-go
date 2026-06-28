package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	expectedConfig := expectedConfig()

	// Set up environment variables for testing
	t.Setenv(envDatabaseURL, expectedConfig.DatabaseURL)
	t.Setenv(envHTTPAddr, expectedConfig.HTTPAddr)

	cfg, err := Load()
	if err != nil {
		t.Fatalf("failed to load configuration: %v", err)
	}

	// Assert the loaded configuration values
	if cfg.DatabaseURL != expectedConfig.DatabaseURL {
		t.Errorf("expected database URL %s, got %s", expectedConfig.DatabaseURL, cfg.DatabaseURL)
	}
	if cfg.HTTPAddr != expectedConfig.HTTPAddr {
		t.Errorf("expected HTTP address %s, got %s", expectedConfig.HTTPAddr, cfg.HTTPAddr)
	}
}

func expectedConfig() Config {
	return Config{
		HTTPAddr:    ":9090",
		DatabaseURL: "postgres://user:pass@localhost:5432/dbname?sslmode=disable",
	}
}
