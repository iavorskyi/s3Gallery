package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *APIServer) listUsers(ctx *gin.Context) {
	userId := ctx.GetInt("userId")

	ctx.IndentedJSON(http.StatusOK, map[string]interface{}{"user_id": userId})

}

func (s *APIServer) getUser(ctx *gin.Context) {

}

func (s *APIServer) updateUser(ctx *gin.Context) {

}

func (s *APIServer) deleteUser(ctx *gin.Context) {

}
