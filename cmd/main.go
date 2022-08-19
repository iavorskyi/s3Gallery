package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/go-pg/pg/v10"
	"github.com/iavorskyi/s3gallery/internal/api"
	"log"
)

var (
	db         *pg.DB
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

	_, err = toml.DecodeFile(configPath, config.DBConfig)
	if err != nil {
		log.Println("Database configs:", err)
	} else {
		log.Println("Using env variables to configure DB")
	}

	s := api.New(config)
	err = s.Start()
	if err == nil {
		log.Fatal("START:", err)
	}

}
