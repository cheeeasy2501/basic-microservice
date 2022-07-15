package config

import (
	"basic-microservice/pkg/database"
	"basic-microservice/pkg/httpserver"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Log  Log
		Http httpserver.HttpServerConfig
		DB   database.DatabaseConfig
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
