package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func UserIdentity(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "empty auth header")
		return
	}

	headersParts := strings.Split(header, " ")
	if len(headersParts) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "invalid auth header")

		return
	}

	userId, err := ParseToken(headersParts[1])
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{"message": "not authorized"})
	}
	ctx.Set("userId", userId)
}
