package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server struct {
		Port string `envconfig:"SERVER_PORT" default:"8080"`
	}

	Database struct {
		Driver      string `envconfig:"DB_DRIVER" default:"sqlite3"`
		DSN         string `envconfig:"DB_DSN" default:"api.db"`
		MaxOpenConn int    `envconfig:"DB_MAX_OPEN_CONNS" default:"10"`
		MaxIdleConn int    `envconfig:"DB_MAX_IDLE_CONNS" default:"5"`
	}

	JWT struct {
		Secret     string        `envconfig:"JWT_SECRET" default:"secret"`
		Expiration time.Duration `envconfig:"JWT_EXPIRATION" default:"1h"`
	}
}

func Load() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
