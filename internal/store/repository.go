package store

import "github.com/iavorskyi/s3gallery/internal/model"

// UserRepository ...
type UserRepository interface {
	Create(user *model.User) error
	FindUserByEmail(email string) (*model.User, error)
	FindByCredential(email, password string) (int, error)
}
