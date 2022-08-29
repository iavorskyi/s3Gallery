package auth

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/iavorskyi/s3gallery/internal/model"
	"github.com/iavorskyi/s3gallery/internal/store"
	database "github.com/iavorskyi/s3gallery/internal/store/sqlStore"

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

func HashPassword(password string) string {
	hash := sha1.New()
	hash.Sum([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte("secretю")))
}

func GenerateToken(user model.User, db *database.Database) (string, error) {
	//var userId int
	//log.Println("TEST:", db)
	//err := db.DBInstance.Model((*s3Gallery.User)(nil)).
	//	Column("id").
	//	Where("email=?", user.Email).
	//	Where("password=?", (user.Password)).
	//	Select(&userId)
	//if err != nil {
	//	return "", err
	//}
	//
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
	//	jwt.StandardClaims{
	//		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
	//		IssuedAt:  time.Now().Unix(),
	//	},
	//	userId,
	//})
	//
	//signedToken, err := token.SignedString([]byte(signingKey))
	//if err != nil {
	//	return "", err
	//}

	return "signedToken", nil
}
