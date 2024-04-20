package users

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
	user handler.User,

) {
	userRoute := []routing.Route{
		{
			Method:     http.MethodPost,
			Path:       "/user",
			Handler:    user.RegisterUser,
			Middleware: []gin.HandlerFunc{},
			Domains:    []string{"v1"},
		}, {
			Method:     http.MethodPut,
			Path:       "/user",
			Handler:    user.UpdateUser,
			Middleware: []gin.HandlerFunc{},
			Domains:    []string{"v1"},
		}, {
			Method:     http.MethodDelete,
			Path:       "/user",
			Handler:    user.DeleteUser,
			Middleware: []gin.HandlerFunc{},
			Domains:    []string{"v1"},
		},
	}
	routing.RegisterRoute(grp, userRoute, log)

}
