package handler

import "github.com/gin-gonic/gin"

type User interface {
	RegisterUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}
