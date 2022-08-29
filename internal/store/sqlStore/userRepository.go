package sqlStore

import (
	"github.com/go-pg/pg/v10"
	"github.com/iavorskyi/s3gallery/internal/model"
	"github.com/iavorskyi/s3gallery/internal/store"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserRepository struct {
	store *Database
}

func (r *UserRepository) Create(user *model.User) error {
	if err := user.BeforeCreate(); err != nil {
		return err
	}

	_, err := r.store.dbInstance.Model(user).Insert()
	if err != nil {
		return err
	}
	return err
}

func (r *UserRepository) FindByCredential(email, password string) (int, error) {
	var userID int
	var user = model.User{}
	err := r.store.dbInstance.Model(&user).
		Where("email=?", email).
		Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return 0, store.ErrRecordNotFound
		}
		return 0, store.ErrRecordNotFound
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println(err)
		return 0, store.ErrWrongPassword
	}
	return userID, err
}

func (r *UserRepository) FindUserByEmail(email string) (*model.User, error) {
	var user = model.User{}
	err := r.store.dbInstance.Model(&user).Where("email = ?", email).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}
