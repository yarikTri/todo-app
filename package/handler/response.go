package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Error describes general error-type
type Error struct {
	Message string `json:""`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)

	c.AbortWithStatusJSON(statusCode, Error{Message: message})
}
