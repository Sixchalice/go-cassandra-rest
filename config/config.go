package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURI string `validate:"required"`
	LogPath     string `validate:"required"`
}

func Load() (*Config, error) {
	// Load .env into OS environment
	_ = godotenv.Load()

	cfg := &Config{
		DatabaseURI: os.Getenv("DATABASE_URI"),
		LogPath:     os.Getenv("LOG_PATH"),
	}

	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return cfg, nil
}
