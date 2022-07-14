package database

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

// postgres config
type DatabaseConfig struct {
	Host     string `env:"DB_HOST" env-default:"localhost"`
	User     string `env:"DB_USER" env-default:"root"`
	Password string `env:"DB_PASSWORD" env-default:"root"`
	DBName   string `env:"DB_NAME" env-required:"true"`
	Port     string `env:"DB_PORT" env-required:"true"`
	SSLMode  string `env:"DB_SSL_MODE" env-default:"disable"`
	// setting *gorm.DB
	SkipDefaultTransaction bool `env:"DB_SKIP_DEFAULT_TRANSACTION" env-default:"true"`
	// setting sql.DB
	SetMaxIdleConns    int `env:"DB_MAX_IDLE_CONNECTIONS" env-default:"3"`
	SetMaxOpenConns    int `env:"DB_MAX_OPEN_CONNECTIONS" env-default:"10"`
	SetConnMaxLifetime int `env:"DB_CONNECTION_MAX_LIFETIME" env-default:"3600"`
	// todo: attemps, timeout
}

func NewConfig() (*DatabaseConfig, error) {
	var err error

	cfg := &DatabaseConfig{}
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (cfg *DatabaseConfig) getDsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)
}
