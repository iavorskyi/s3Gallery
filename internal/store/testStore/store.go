package testStore

import (
	"github.com/iavorskyi/s3gallery/internal/model"
	"github.com/iavorskyi/s3gallery/internal/store"
)

type Database struct {
	userRepository *UserRepository
}

func New() *Database {
	return &Database{}
}

func (d *Database) User() store.UserRepository {
	if d.userRepository != nil {
		return d.userRepository
	}
	d.userRepository = &UserRepository{
		store: d,
		users: make(map[string]*model.User),
	}
	return d.userRepository
}
