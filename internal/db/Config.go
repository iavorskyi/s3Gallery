package database

import (
	"log"
	"os"
)

type Config struct {
	ConnectionString string `toml:"connect_str"`
	User             string `toml:"user"`
	Password         string `toml:"password"`
	Name             string `toml:"name"`
}

func NewConfig() *Config {
	config := Config{}
	config.User = os.Getenv("DB_USER")
	config.Name = os.Getenv("DB_NAME")
	config.Password = os.Getenv("DB_PASSWORD")
	config.ConnectionString = os.Getenv("DB_CONN_STR")
	log.Println("ENV_CONF:", config)
	return &config
}
