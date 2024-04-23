package handler

import "github.com/gin-gonic/gin"

type User interface {
	RegisterUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetUsers(c *gin.Context)
}
type PaymentRule interface {
	CreatePaymentRule(c *gin.Context)
	UpdatepaymentRule(c *gin.Context)
	GetPaymentRule(c *gin.Context)
	DeletepaymentRule(c *gin.Context)
}
