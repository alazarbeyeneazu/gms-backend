package paymentrule

import (
	"net/http"

	"github.com/alazarbeyeneazu/gms-backend/internal/glue/routing"
	"github.com/alazarbeyeneazu/gms-backend/internal/handler"
	"github.com/alazarbeyeneazu/gms-backend/internal/handler/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Init(
	grp *gin.RouterGroup,
	log zap.Logger,
	user handler.PaymentRule,

) {
	paymentRuleRoute := []routing.Route{
		{
			Method:     http.MethodPost,
			Path:       "/payment/rule",
			Handler:    user.CreatePaymentRule,
			Middleware: []gin.HandlerFunc{},
			Domains:    []string{"v1"},
		}, {
			Method:     http.MethodPut,
			Path:       "/payment/rule",
			Handler:    user.UpdatepaymentRule,
			Middleware: []gin.HandlerFunc{},
			Domains:    []string{"v1"},
		}, {
			Method:  http.MethodGet,
			Path:    "/payment/rule",
			Handler: user.GetPaymentRule,
			Middleware: []gin.HandlerFunc{
				middleware.Auth(),
			},
			Domains: []string{"v1"},
		}, {
			Method:     http.MethodDelete,
			Path:       "/payment/rule",
			Handler:    user.DeletepaymentRule,
			Middleware: []gin.HandlerFunc{},
			Domains:    []string{"v1"},
		},
	}
	routing.RegisterRoute(grp, paymentRuleRoute, log)

}
