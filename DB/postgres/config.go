package postgres

import (
	"github.com/go-pg/pg/v10"
)

type BDConfig struct {
	ConnectionString string
	User             string
	Password         string
	Name             string
}

func GetDb(config BDConfig) *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     config.ConnectionString,
		User:     config.User,
		Password: config.Password,
		Database: config.Name,
	})
}
