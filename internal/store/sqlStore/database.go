package sqlStore

import (
	"github.com/go-pg/pg/v10"
	"github.com/iavorskyi/s3gallery/internal/store"
)

type Database struct {
	dbInstance     *pg.DB
	userRepository *UserRepository
}

func New(db *pg.DB) *Database {
	return &Database{dbInstance: db}
}

func (d *Database) User() store.UserRepository {
	if d.userRepository != nil {
		return d.userRepository
	}
	d.userRepository = &UserRepository{
		store: d,
	}
	return d.userRepository
}
