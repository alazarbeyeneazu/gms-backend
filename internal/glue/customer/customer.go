package customer

import (
	"net/http"

	"github.com/alazarbeyeneazu/gms-backend/internal/glue/routing"
	"github.com/alazarbeyeneazu/gms-backend/internal/handler"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Init(
	grp *gin.RouterGroup,
	log zap.Logger,
	user handler.Customer,

) {
	paymentRuleRoute := []routing.Route{
		{
			Method:     http.MethodPost,
			Path:       "/customer",
			Handler:    user.RegisterCustomer,
			Middleware: []gin.HandlerFunc{},
			Domains:    []string{"v1"},
		}, {
			Method:     http.MethodGet,
			Path:       "/customer",
			Handler:    user.GetCustomer,
			Middleware: []gin.HandlerFunc{},
			Domains:    []string{"v1"},
		}, {
			Method:     http.MethodPut,
			Path:       "/customer",
			Handler:    user.UpdateCustomer,
			Middleware: []gin.HandlerFunc{},
			Domains:    []string{"v1"},
		}, {
			Method:     http.MethodDelete,
			Path:       "/customer",
			Handler:    user.DeleteCustomer,
			Middleware: []gin.HandlerFunc{},
			Domains:    []string{"v1"},
		},
	}
	routing.RegisterRoute(grp, paymentRuleRoute, log)

}
