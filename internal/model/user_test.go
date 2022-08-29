package model_test

import (
	"github.com/iavorskyi/s3gallery/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_BeforeCreate(t *testing.T) {
	user := model.TestUser(t)
	password := user.Password
	assert.NoError(t, user.BeforeCreate())
	assert.NotNil(t, user.Password)
	assert.NotEqual(t, user.Password, password)
}

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		user    func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			user: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty email",
			user: func() *model.User {
				user := model.TestUser(t)
				user.Email = ""
				return user
			},
			isValid: false,
		},
		{
			name: "not email",
			user: func() *model.User {
				user := model.TestUser(t)
				user.Email = "notEmail"
				return user
			},
			isValid: false,
		},
		{
			name: "empty password",
			user: func() *model.User {
				user := model.TestUser(t)
				user.Password = ""
				return user
			},
			isValid: false,
		},
		{
			name: "short password",
			user: func() *model.User {
				user := model.TestUser(t)
				user.Password = "1234"
				return user
			},
			isValid: false,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			if c.isValid {
				assert.NoError(t, c.user().Validate())
			} else {
				assert.Error(t, c.user().Validate())
			}
		})
	}
}
