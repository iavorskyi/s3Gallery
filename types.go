package s3Gallery

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

type User struct {
	tableName struct{} `pg:"users"`
	ID        string   `json:"id"`
	Email     string   `json:"email" binding:"required"`
	FirstName string   `json:"first_name,omitempty"`
	LastName  string   `json:"last_name,omitempty"`
	Password  string   `json:"password" binding:"required"`
}

type Album struct {
	Name         string
	CreationTime time.Time
}

type Item struct {
	Type         string    `json:"type"`
	LastModified time.Time `json:"last-modified"`
	Album        string    `json:"album"`
	Name         string    `json:"name"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
