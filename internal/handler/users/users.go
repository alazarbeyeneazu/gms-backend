package goenergy

import (
	"net/http"

	"github.com/Adamant-Investment-PLC/Backend/internal/constants/model/response"
	"github.com/Adamant-Investment-PLC/Backend/internal/handler"
	"github.com/Adamant-Investment-PLC/Backend/internal/module"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type user struct {
	UserModule module.User
	Logger     zap.Logger
}

func Init(userModule module.User, log zap.Logger) handler.User {
	return &user{
		UserModule: userModule,
		Logger:     log,
	}
}
func (u *user) CreateUser(c *gin.Context) {

	response.SendSuccessResponse(c, http.StatusCreated, "")
	return
}
