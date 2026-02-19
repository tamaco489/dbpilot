package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Driver          string `mapstructure:"driver"`
	DSN             string `mapstructure:"dsn"`
	ReadOnly        bool   `mapstructure:"read_only"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	if err := cfg.validate(); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *Config) validate() error {
	if c.Driver != "mysql" && c.Driver != "postgres" {
		return fmt.Errorf("invalid driver %q: must be \"mysql\" or \"postgres\"", c.Driver)
	}
	if c.DSN == "" {
		return fmt.Errorf("dsn is required")
	}
	return nil
}
