package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	Request(c *gin.Context)
}