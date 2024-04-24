package initiator

import (
	"github.com/gin-gonic/gin"

	"github.com/alazarbeyeneazu/gms-backend/internal/glue/customer"
	"github.com/alazarbeyeneazu/gms-backend/internal/glue/paymentrule"
	"github.com/alazarbeyeneazu/gms-backend/internal/glue/users"
	"go.uber.org/zap"
)

func InitRouting(
	grp *gin.RouterGroup,
	log zap.Logger,
	handler Handler,
) {
	users.Init(grp, log, handler.User)
	paymentrule.Init(grp, log, handler.PaymentRule)
	customer.Init(grp, log, handler.customer)

}
