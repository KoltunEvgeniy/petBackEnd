package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Message string
}

func NewErrorResponce(c *gin.Context, statusCode int, msg string) {
	logrus.Error(msg)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{Message: msg})
}
