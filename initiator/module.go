package initiator

import (
	"github.com/alazarbeyeneazu/gms-backend/internal/module"
	"github.com/alazarbeyeneazu/gms-backend/internal/module/auth"
	"github.com/alazarbeyeneazu/gms-backend/internal/module/customer"
	"github.com/alazarbeyeneazu/gms-backend/internal/module/paymentrule"
	"github.com/alazarbeyeneazu/gms-backend/internal/module/users"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Module struct {
	User        module.User
	PaymentRule module.PaymentRule
	customer    module.Customer
	auth        module.Auth
}

func InitModule(log *zap.Logger, persistenceDb Persistence) Module {
	return Module{
		User:        users.Init(persistenceDb.user, log),
		PaymentRule: paymentrule.Init(persistenceDb.paymentRule, log),
		customer:    customer.Init(persistenceDb.customer, persistenceDb.paymentRule, persistenceDb.user, log),
		auth:        auth.Init(persistenceDb.user, persistenceDb.auth, viper.GetString("jwt_key"), log),
	}
}
