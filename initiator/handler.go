package initiator

import (
	"github.com/alazarbeyeneazu/gms-backend/internal/handler"
	"github.com/alazarbeyeneazu/gms-backend/internal/handler/customer"
	"github.com/alazarbeyeneazu/gms-backend/internal/handler/paymentrule"
	"github.com/alazarbeyeneazu/gms-backend/internal/handler/users"
	"go.uber.org/zap"
)

type Handler struct {
	User        handler.User
	PaymentRule handler.PaymentRule
	customer    handler.Customer
}

func InitHandler(module Module, log zap.Logger) Handler {
	return Handler{
		User:        users.Init(module.User, log),
		PaymentRule: paymentrule.Init(module.PaymentRule, log),
		customer:    customer.Init(module.customer, log),
	}
}
