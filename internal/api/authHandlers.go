package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iavorskyi/s3gallery/Services/auth"
	"github.com/iavorskyi/s3gallery/internal/model"
	"github.com/iavorskyi/s3gallery/s3Gallery"
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

//
//func (s *Server) signIn(ctx *gin.Context) {
//	var user model.User
//	err := ctx.BindJSON(&user)
//	if err != nil {
//		ctx.String(http.StatusBadRequest, err.Error())
//		return
//	}
//	token, err := auth.GenerateToken(user, s.store)
//	if err != nil {
//		log.Println("Failed to sign in", user.Email, err)
//		ctx.String(http.StatusInternalServerError, err.Error())
//		s3Gallery.NewErrorResponse(ctx, http.StatusInternalServerError, "Failed to sign in"+user.Email+err.Error())
//		return
//	}
//
//	ctx.IndentedJSON(http.StatusOK, map[string]interface{}{"token": token, "user": user.Email})
//
//}
