package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iavorskyi/s3gallery/Services/auth"
	"github.com/iavorskyi/s3gallery/s3Gallery"
	"log"
	"net/http"
)

func (s *APIServer) signUp(ctx *gin.Context) {
	var newUser s3Gallery.User
	err := ctx.BindJSON(&newUser)
	if err != nil {
		s3Gallery.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	createdUser, err := auth.CreateUser(newUser, s.store)
	if err != nil {
		s3Gallery.NewErrorResponse(ctx, http.StatusInternalServerError, "Failed to create user"+newUser.Email+err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, createdUser)
}

func (s *APIServer) signIn(ctx *gin.Context) {
	var user s3Gallery.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	token, err := auth.GenerateToken(user, s.store)
	if err != nil {
		log.Println("Failed to sign in", user.Email, err)
		ctx.String(http.StatusInternalServerError, err.Error())
		s3Gallery.NewErrorResponse(ctx, http.StatusInternalServerError, "Failed to sign in"+user.Email+err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, map[string]interface{}{"token": token, "user": user.Email})

}
