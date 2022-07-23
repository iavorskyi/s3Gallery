package auth

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-pg/pg/v10"
	s3Gallery "github.com/iavorskyi/s3gallery"
	"time"
)

const (
	signingKey = "someSecretKey"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func CreateUser(user s3Gallery.User, db *pg.DB) (s3Gallery.User, error) {
	var createdUser s3Gallery.User

	user.Password = hashPassword(user.Password)
	if !validateEmail(user.Email) {
		return createdUser, errors.New("email is not valid")
	}
	_, err := db.Model(&user).Insert()
	if err != nil {
		return createdUser, err
	}

	err = db.Model(&createdUser).Where("email = ?", user.Email).Select()
	if err != nil {
		return createdUser, err
	}
	return createdUser, nil
}

func hashPassword(password string) string {
	hash := sha1.New()
	hash.Sum([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte("secret")))
}

func validateEmail(email string) bool {
	return true
}

func GenerateToken(user s3Gallery.User, db *pg.DB) (string, error) {
	var userId int
	err := db.Model((*s3Gallery.User)(nil)).
		Column("id").
		Where("email=?", user.Email).
		Where("password=?", hashPassword(user.Password)).
		Select(&userId)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})

	signedToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
