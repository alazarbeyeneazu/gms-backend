package auth

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
	user handler.Auth,

) {
	userRoute := []routing.Route{
		{
			Method:     http.MethodPost,
			Path:       "/user/add/login",
			Handler:    user.RegisterUserAuth,
			Middleware: []gin.HandlerFunc{},
			Domains:    []string{"v1"},
		}, {
			Method:     http.MethodPost,
			Path:       "/user/login",
			Handler:    user.Login,
			Middleware: []gin.HandlerFunc{},
			Domains:    []string{"v1"},
		},
	}
	routing.RegisterRoute(grp, userRoute, log)

}
