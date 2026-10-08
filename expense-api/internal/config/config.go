package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
)

// Load variables
type Config struct {
	DatabaseURL string `env:"DATABASE_URL"`
	Port        string `env:"PORT" envDefault:"8080"`

	DBMaxConns        int32         `env:"DB_MAX_CONNS" envDefault:"10"`
	DBMinConns        int32         `env:"DB_MIN_CONNS" envDefault:"2"`
	DBMaxConnLifetime time.Duration `env:"DB_MAX_CONN_LIFETIME" envDefault:"30m"`
	DBMaxConnIdleTime time.Duration `env:"DB_MAX_CONN_IDLE_TIME" envDefault:"5m"`
}

// Parse config
func Load() (*Config, error) {
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return &cfg, nil
}
