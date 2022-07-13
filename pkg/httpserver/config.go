package httpserver

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type IHttpServerConfig interface {
	GetAddr() string
}

type HttpServerConfig struct {
	Host         string `env:"HTTP_HOST" env-default:""`
	Port         string `env:"HTTP_PORT" env-default:"80"`
	ReadTimeout  int    `env:"HTTP_READ_TIMEOUT" env-default:"5"`
	WriteTimeout int    `env:"HTTP_WRITE_TIMEOUT" env-default:"5"`
}

func (c *HttpServerConfig) GetAddr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func NewConfig() (*HttpServerConfig, error) {
	var err error
	cfg := &HttpServerConfig{}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
