package api

import (
	"context"
	"github.com/go-pg/pg/v10"
	awsStore "github.com/iavorskyi/s3gallery/internal/store/awsS3"
	"github.com/iavorskyi/s3gallery/internal/store/sqlStore"
	"log"
	"net/http"
)

func Start(config *Config) error {
	db, err := NewDB(config.DBConnectStr, config.DBUser, config.DBName, config.DBPassword)
	if err != nil {
		return err
	}
	defer db.Close()

	s3, err := awsStore.GetClient()
	if err != nil {
		return err
	}
	manager, err := awsStore.GetManager()
	if err != nil {
		return err
	}

	store := sqlStore.New(db)

	s3store := awsStore.New(s3, manager)
	srv := newServer(store, s3store)
	log.Println("Starting on", config.BindAddr, "port")

	return http.ListenAndServe(config.BindAddr, srv)
}

func NewDB(connString, dbUser, dbName, dbPassword string) (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		Addr:     connString,
		User:     dbUser,
		Password: dbPassword,
		Database: dbName,
	})

	if err := db.Ping(context.Background()); err != nil {
		return nil, err
	}
	return db, nil
}
