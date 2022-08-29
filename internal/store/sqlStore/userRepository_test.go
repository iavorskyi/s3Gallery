package sqlStore_test

import (
	"github.com/iavorskyi/s3gallery/internal/model"
	"github.com/iavorskyi/s3gallery/internal/store"
	database "github.com/iavorskyi/s3gallery/internal/store/sqlStore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := database.TestDB(t, dbConnectionString, dbUser, dbName, dbPassword)
	defer teardown("users")
	store := database.New(db)

	user := model.TestUser(t)
	err := store.User().Create(user)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_FindUserByEmail(t *testing.T) {
	db, teardown := database.TestDB(t, dbConnectionString, dbUser, dbName, dbPassword)
	defer teardown("users")

	s := database.New(db)

	email := "user@example.com"

	_, err := s.User().FindUserByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	user := model.TestUser(t)
	s.User().Create(user)
	createdUser, err := s.User().FindUserByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
}

func TestUserRepository_FindByCredential(t *testing.T) {
	db, teardown := database.TestDB(t, dbConnectionString, dbUser, dbName, dbPassword)
	defer teardown("users")

	s := database.New(db)

	// with wrong email and password
	wrongEmail := "wrong_user@example.com"
	wrongPassword := "wrong"
	_, err := s.User().FindByCredential(wrongEmail, wrongPassword)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	user := model.TestUser(t)
	s.User().Create(user)

	// with wrong password
	_, err = s.User().FindByCredential(user.Email, wrongPassword)
	assert.EqualError(t, err, store.ErrWrongPassword.Error())

	// success test
	user = model.TestUser(t)
	gottenUser, err := s.User().FindByCredential(user.Email, user.Password)
	assert.NoError(t, err)
	assert.NotNil(t, gottenUser)
}
