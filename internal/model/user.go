package model

import (
	"github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

const SigningKey = "someSecretKey"

type User struct {
	tableName struct{} `pg:"users"`
	ID        int      `json:"id"`
	Email     string   `json:"email" binding:"required"`
	FirstName string   `json:"first_name,omitempty"`
	LastName  string   `json:"last_name,omitempty"`
	Password  string   `json:"password" binding:"required"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(5, 16)),
	)
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (u *User) BeforeCreate() error {
	var err error
	u.Password, err = HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
