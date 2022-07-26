package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/iavorskyi/s3gallery/internal/store"
	"github.com/sirupsen/logrus"
	"net/http"
)

type server struct {
	router       *gin.Engine
	logger       *logrus.Logger
	store        store.Store
	s3store      store.S3Store
	sessionStore sessions.Store
}

func newServer(store store.Store, s3store store.S3Store, sessionStore sessions.Store) *server {
	s := &server{
		router:       gin.Default(),
		logger:       logrus.New(),
		store:        store,
		s3store:      s3store,
		sessionStore: sessionStore,
	}
	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(sessions.Sessions("mysession", s.sessionStore))

	s.router.POST("/sign-up", s.signUp)
	s.router.POST("/sign-in", s.signIn)
	s.router.GET("/sign-out", s.signOut)

	api := s.router.Group("/api", s.userIdentity)
	//api.GET("/", s.apiPing)
	//
	//users := api.Group("/users")
	//users.GET("/", s.listUsers)
	//users.GET("/:id", s.getUser)
	//users.PUT("/:id", s.updateUser)
	//users.DELETE("/:id", s.deleteUser)
	//
	albums := api.Group("/albums")
	albums.GET("/", s.listAlbums)
	//albums.GET("/:albumId", s.getAlbum)
	//albums.POST("/", s.createAlbum)
	//
	items := albums.Group(":albumId/items")
	items.POST("/", s.uploadItem)
	items.GET("/", s.listItems)
	items.GET("/:id", s.getItem)
	//items.PUT("/:id", s.updateItem)
	items.DELETE("/:id", s.deleteItem)
}
