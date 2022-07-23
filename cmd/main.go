package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/iavorskyi/s3gallery/DB/postgres"
	"github.com/iavorskyi/s3gallery/Services/auth"
)

var db *pg.DB

func main() {
	dbConfig := postgres.BDConfig{
		Name:             "S3Gallery",
		User:             "admin",
		Password:         "uICklV9e6m2FZ40yRUTfA7gw53OoED81",
		ConnectionString: "185.226.42.227:5432",
	}
	db = postgres.GetDb(dbConfig)

	router := gin.Default()

	router.POST("/sign-up", signUp)
	router.POST("/sign-in", signIn)

	api := router.Group("/api", auth.UserIdentity)

	users := api.Group("/users")
	users.GET("/", listUsers)
	users.GET("/:id", getUser)
	users.PUT("/:id", updateUser)
	users.DELETE("/:id", deleteUser)

	albums := api.Group("/albums")
	albums.GET("/", listAlbums)
	albums.GET("/:albumId", getAlbum)
	albums.POST("/", createAlbum)

	items := albums.Group(":albumId/items")
	items.POST("/", uploadItem)
	items.GET("/", listItems)
	items.GET("/:id", getItem)
	items.PUT("/:id", updateItem)
	items.DELETE("/:id", deleteItem)

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
