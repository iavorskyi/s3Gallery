package testStore

import (
	"errors"
	"github.com/iavorskyi/s3gallery/internal/model"
	"github.com/iavorskyi/s3gallery/internal/store"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserRepository struct {
	store *Database
	users map[string]*model.User
}

// Create ...
func (r *UserRepository) Create(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	if err := user.BeforeCreate(); err != nil {
		return err
	}

	r.users[user.Email] = user
	user.ID = len(r.users)

	return nil
}

// FindUserByEmail ...
func (r *UserRepository) FindUserByEmail(email string) (*model.User, error) {
	if user, ok := r.users[email]; !ok {
		return nil, store.ErrRecordNotFound
	} else {
		return user, nil
	}
}

// FindByCredential ...
func (r *UserRepository) FindByCredential(email string, password string) (int, error) {
	if user, ok := r.users[email]; !ok {
		return 0, store.ErrRecordNotFound
	} else {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			log.Println(err)
			return 0, errors.New("wrong password")
		}
		return user.ID, nil
	}
}
