package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/iavorskyi/s3gallery/Services/auth"
	"github.com/iavorskyi/s3gallery/internal/model"
	"github.com/iavorskyi/s3gallery/s3Gallery"
	"log"
	"net/http"
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
		s3Gallery.NewErrorResponse(ctx, http.StatusInternalServerError, "Failed to sign in"+user.Email+err.Error())
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

	ctx.IndentedJSON(http.StatusOK, map[string]interface{}{"message": "logged out"})
}
