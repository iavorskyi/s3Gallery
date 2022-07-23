package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func listUsers(ctx *gin.Context) {
	userId := ctx.GetInt("userId")

	ctx.IndentedJSON(http.StatusOK, map[string]interface{}{"user_id": userId})

}

func getUser(ctx *gin.Context) {

}

func updateUser(ctx *gin.Context) {

}

func deleteUser(ctx *gin.Context) {

}
