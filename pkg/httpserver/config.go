package httpserver

type HttpServerConfig struct {
	Host string `env:"HTTP_HOST" env-default:""`
	Port string `env:"HTTP_PORT" env-default:"80"`
}
