package response

import (
	"github.com/gin-gonic/gin"
	"github.com/inflexjs/crud-backend/internal/models"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

type GetAllPostsResponse struct {
	Data []models.Post `json:"data"`
}

func NewError(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{message})
}
