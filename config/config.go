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
	Port        int    `validate:"min=1,max=65535" env:"PORT" envDefault:"8080"`
}

func Load() (*Config, error) {
	// Load .env into OS environment
	_ = godotenv.Load()

	port := 8080
	if os.Getenv("PORT") != "" {
		fmt.Sscanf(os.Getenv("PORT"), "%d", &port)
	}

	cfg := &Config{
		DatabaseURI: os.Getenv("DATABASE_URI"),
		LogPath:     os.Getenv("LOG_PATH"),
		Port:        port,
	}

	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return cfg, nil
}
