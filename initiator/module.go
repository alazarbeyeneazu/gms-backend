package initiator

import (
	"github.com/alazarbeyeneazu/gms-backend/internal/module"
	"github.com/alazarbeyeneazu/gms-backend/internal/module/users"
	"go.uber.org/zap"
)

type Module struct {
	User module.User
}

func InitModule(log *zap.Logger, persistenceDb Persistence) Module {
	return Module{
		User: users.Init(persistenceDb.user, log),
	}
}
