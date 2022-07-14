package exception

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Exception struct {
	Message string `json:"message"`
}

func NewError(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, Exception{message})
}
