package initiator

import (
	"go.uber.org/zap"
)

type Module struct {
}

func InitModule(log *zap.Logger, persistenceDb Persistence) Module {
	return Module{}
}
