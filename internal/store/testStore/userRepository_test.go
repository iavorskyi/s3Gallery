package testStore_test

import (
	"github.com/iavorskyi/s3gallery/internal/model"
	"github.com/iavorskyi/s3gallery/internal/store"
	"github.com/iavorskyi/s3gallery/internal/store/testStore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := testStore.New()

	user := model.TestUser(t)
	assert.NoError(t, s.User().Create(user))
	assert.NotNil(t, user)
}

func TestUserRepository_FindUserByEmail(t *testing.T) {
	s := testStore.New()
	email := "user@example.com"
	_, err := s.User().FindUserByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	user := model.TestUser(t)
	s.User().Create(user)
	user, err = s.User().FindUserByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_FindByCredential(t *testing.T) {
	s := testStore.New()

	// with wrong email and password
	wrongEmail := "wrong_user@example.com"
	wrongPassword := "wrong"
	_, err := s.User().FindByCredential(wrongEmail, wrongPassword)
	assert.Error(t, err)

	user := model.TestUser(t)
	s.User().Create(user)

	// with wrong password
	_, err = s.User().FindByCredential(user.Email, wrongPassword)
	assert.Error(t, err)

	// success test
	user = model.TestUser(t)
	gottenUser, err := s.User().FindByCredential(user.Email, user.Password)
	assert.NoError(t, err)
	assert.NotNil(t, gottenUser)
}
