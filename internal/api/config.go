package api

import "github.com/iavorskyi/s3gallery/internal/db"

// Config for API server
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	DBConfig *database.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8000",
		LogLevel: "debug",
		DBConfig: database.NewConfig(),
	}
}
