package initiator

import (
	"github.com/gin-gonic/gin"

	"github.com/alazarbeyeneazu/gms-backend/internal/glue/users"
	"go.uber.org/zap"
)

func InitRouting(
	grp *gin.RouterGroup,
	log zap.Logger,
	handler Handler,
) {
	users.Init(grp, log, handler.User)
}
