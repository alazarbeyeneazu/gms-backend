package initiator

import (
	"github.com/alazarbeyeneazu/gms-backend/internal/module"
	"github.com/alazarbeyeneazu/gms-backend/internal/module/paymentrule"
	"github.com/alazarbeyeneazu/gms-backend/internal/module/users"
	"go.uber.org/zap"
)

type Module struct {
	User        module.User
	PaymentRule module.PaymentRule
}

func InitModule(log *zap.Logger, persistenceDb Persistence) Module {
	return Module{
		User:        users.Init(persistenceDb.user, log),
		PaymentRule: paymentrule.Init(persistenceDb.paymentRule, log),
	}
}
