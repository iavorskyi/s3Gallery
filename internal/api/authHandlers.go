package api

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/iavorskyi/s3gallery/Services/auth"
	"github.com/iavorskyi/s3gallery/internal/model"
	"github.com/iavorskyi/s3gallery/s3Gallery"
	"log"
	"net/http"
	"strings"
)

func (s *server) signUp(ctx *gin.Context) {
	var newUser model.User
	err := ctx.BindJSON(&newUser)
	if err != nil {
		s3Gallery.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	createdUser, err := auth.CreateUser(newUser, s.store)
	if err != nil {
		s3Gallery.NewErrorResponse(ctx, http.StatusInternalServerError, "Failed to create user"+newUser.Email+err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusCreated, createdUser)
}

func (s *server) signIn(ctx *gin.Context) {
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	token, err := auth.GenerateToken(user, s.store)
	if err != nil {
		log.Println("Failed to sign in", user.Email, err)
		s3Gallery.NewErrorResponse(ctx, http.StatusInternalServerError, "Failed to sign in "+user.Email+err.Error())
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	session := sessions.Default(ctx)
	session.Set("user", user.Email)
	if err = session.Save(); err != nil {
		s3Gallery.NewErrorResponse(ctx, http.StatusInternalServerError, "Failed to save session "+err.Error())
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.IndentedJSON(http.StatusOK, map[string]interface{}{"token": token, "user": user.Email})

}

func (s *server) signOut(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	err := session.Save()
	if err != nil {
		ctx.IndentedJSON(http.StatusOK, map[string]interface{}{"message": err.Error()})
	}
	ctx.IndentedJSON(http.StatusOK, map[string]interface{}{"message": "logged out"})
}

func ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &model.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(model.SigningKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*model.TokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (s *server) userIdentity(ctx *gin.Context) {
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
	//session := sessions.Default(ctx)
	//email := session.Get("user")
	//if email == nil {
	//	ctx.JSON(http.StatusNotFound, gin.H{
	//		"message": "unauthorized",
	//	})
	//	ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{"message": "not authorized"})
	//
	//}
	ctx.Set("userId", userId)
}
