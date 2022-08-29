package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/iavorskyi/s3gallery/internal/api"
	"log"
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

	if err := api.Start(config); err != nil {
		log.Fatal(err)
	}
}
