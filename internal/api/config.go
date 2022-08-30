package api

// Config for API server
type Config struct {
	BindAddr     string `toml:"bind_addr"`
	LogLevel     string `toml:"log_level"`
	DBConnectStr string `toml:"connect_str"`
	DBUser       string `toml:"user"`
	DBName       string `toml:"name"`
	DBPassword   string `toml:"password"`
	S3Endpoint   string `toml:"s3endpoint"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8000",
		LogLevel: "debug",
	}
}
