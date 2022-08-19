package database

import (
	"context"
	"github.com/go-pg/pg/v10"
)

type Database struct {
	config     *Config
	DBInstance *pg.DB
}

func New(config *Config) *Database {
	dbInstance := pg.DB{}
	return &Database{config: config, DBInstance: &dbInstance}
}

func (db *Database) Open() error {
	db.DBInstance = pg.Connect(&pg.Options{
		Addr:     db.config.ConnectionString,
		User:     db.config.User,
		Password: db.config.Password,
		Database: db.config.Name,
	})
	if err := db.DBInstance.Ping(context.Background()); err != nil {
		return err
	}
	return nil
}

func (db *Database) Close() error {
	if err := db.DBInstance.Close(); err != nil {
		return err
	}
	return nil
}
