package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iavorskyi/s3gallery/Services/auth"
	"github.com/iavorskyi/s3gallery/internal/db"
	"github.com/sirupsen/logrus"
	"net/http"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *gin.Engine
	store  *database.Database
}

// New default API server
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: gin.Default(),
	}
}

// Start API server
func (s *APIServer) Start() error {
	if err := s.ConfigureLogger(); err != nil {
		return err
	}
	s.logger.Info("API server is starting...")

	s.ConfigureRouter()
	if err := s.ConfigureDB(); err != nil {
		return err
	}

	return s.router.Run(s.config.BindAddr)
}

func (s *APIServer) ConfigureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) ConfigureRouter() {
	s.router.POST("/sign-up", s.signUp)
	s.router.POST("/sign-in", s.signIn)

	api := s.router.Group("/api")
	api.GET("/", s.apiPing)

	users := api.Group("/users")
	users.GET("/", s.listUsers)
	users.GET("/:id", s.getUser)
	users.PUT("/:id", s.updateUser)
	users.DELETE("/:id", s.deleteUser)

	albums := api.Group("/albums", auth.UserIdentity)
	albums.GET("/", s.listAlbums)
	albums.GET("/:albumId", s.getAlbum)
	albums.POST("/", s.createAlbum)

	items := albums.Group(":albumId/items")
	items.POST("/", s.uploadItem)
	items.GET("/", s.listItems)
	items.GET("/:id", s.getItem)
	items.PUT("/:id", s.updateItem)
	items.DELETE("/:id", s.deleteItem)
}

func (s *APIServer) ConfigureDB() error {
	database := database.New(s.config.DBConfig)
	if err := database.Open(); err != nil {
		return err
	}
	s.store = database
	return nil
}

func (s *APIServer) apiPing(ctx *gin.Context) {
	ctx.String(http.StatusOK, "PONG")
}

//func (s *APIServer) ConfigureDB() {
//	dbConfig := postgres.BDConfig{
//		Name:             "S3Gallery",
//		User:             "admin",
//		Password:         "uICklV9e6m2FZ40yRUTfA7gw53OoED81",
//		ConnectionString: "185.226.42.227:5432",
//	}
//	s.db = postgres.GetDb(dbConfig)
//}
