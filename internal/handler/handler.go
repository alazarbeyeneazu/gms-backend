package handler

import "github.com/gin-gonic/gin"

type User interface {
	CreateUser(c *gin.Context)
}
