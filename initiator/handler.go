package initiator

import (
	"go.uber.org/zap"
)

type Handler struct {
}

func InitHandler(module Module, log zap.Logger) Handler {
	return Handler{}
}
