package s3Gallery

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

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
