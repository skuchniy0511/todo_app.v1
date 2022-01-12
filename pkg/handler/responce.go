package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponce struct {
	Message string ` json:"message"`
}

type StatusResponce struct {
	Status string `json:"status"`
}

func NewErrorResponce(c *gin.Context, statuscode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statuscode, errorResponce{message})
}
