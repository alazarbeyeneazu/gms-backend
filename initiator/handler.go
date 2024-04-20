package initiator

import (
	"github.com/alazarbeyeneazu/gms-backend/internal/handler"
	"github.com/alazarbeyeneazu/gms-backend/internal/handler/users"
	"go.uber.org/zap"
)

type Handler struct {
	User handler.User
}

func InitHandler(module Module, log zap.Logger) Handler {
	return Handler{
		User: users.Init(module.User, log),
	}
}
