package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/iavorskyi/s3gallery/internal/model"
	"github.com/iavorskyi/s3gallery/internal/store"
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

func CreateUser(user model.User, db store.Store) (*model.User, error) {
	err := db.User().Create(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GenerateToken(user model.User, db store.Store) (string, error) {
	userID, err := db.User().FindByCredential(user.Email, user.Password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userID,
	})

	signedToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
