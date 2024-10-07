package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -
	Config struct {
		HTTP
		DB
	}
	// HTTP -
	HTTP struct {
		Port string `env-default:"8080" env:"HTTP_PORT"`
	}
	// DB -
	DB struct {
		URL    string `env-required:"true" env:"DB_URL"`
		Driver string `env-default:"postgres" env:"DB_DRIVER"`
	}
)

// NewConfig returns app config
func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
