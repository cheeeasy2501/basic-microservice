package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Log Log
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"  env:"LOG_LEVEL"`
	}
)

func NewConfig() (*Config, error) {
	var err error
	cfg := &Config{}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
