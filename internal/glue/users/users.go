package users

import (
	"net/http"

	"github.com/Adamant-Investment-PLC/Backend/internal/glue/routing"
	"github.com/Adamant-Investment-PLC/Backend/internal/handler"

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
			Handler:    user.CreateUser,
			Middleware: []gin.HandlerFunc{},
			Domains:    []string{"v1"},
		},
	}
	routing.RegisterRoute(grp, userRoute, log)

}
