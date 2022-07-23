package main

import (
	"github.com/gin-gonic/gin"
	S3Gallery "github.com/iavorskyi/s3gallery"
	"github.com/iavorskyi/s3gallery/Services/auth"
	"log"
	"net/http"
)

func signUp(ctx *gin.Context) {
	var newUser S3Gallery.User
	err := ctx.BindJSON(&newUser)
	if err != nil {
		S3Gallery.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	createdUser, err := auth.CreateUser(newUser, db)
	if err != nil {
		S3Gallery.NewErrorResponse(ctx, http.StatusInternalServerError, "Failed to create user"+newUser.Email+err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, createdUser)
}

func signIn(ctx *gin.Context) {
	var user S3Gallery.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	token, err := auth.GenerateToken(user, db)
	if err != nil {
		log.Println("Failed to sign in", user.Email, err)
		ctx.String(http.StatusInternalServerError, err.Error())
		S3Gallery.NewErrorResponse(ctx, http.StatusInternalServerError, "Failed to sign in"+user.Email+err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, map[string]interface{}{"token": token})

}
