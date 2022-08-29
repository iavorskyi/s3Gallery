package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/iavorskyi/s3gallery/internal/api"
	"log"
	"os"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "configpath", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("app configs:", err)
	}
	config.DBName = os.Getenv("DB_NAME")
	config.DBUser = os.Getenv("DB_USER")
	config.DBPassword = os.Getenv("DB_PASSWORD")
	config.DBConnectStr = os.Getenv("DB_CONN_STR")

	if err := api.Start(config); err != nil {
		log.Fatal(err)
	}
}
