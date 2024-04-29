package auth

import (
	"net/http"

	"github.com/alazarbeyeneazu/gms-backend/internal/constants/errors"
	"github.com/alazarbeyeneazu/gms-backend/internal/constants/model/dto"
	"github.com/alazarbeyeneazu/gms-backend/internal/constants/model/response"
	"github.com/alazarbeyeneazu/gms-backend/internal/handler"
	"github.com/alazarbeyeneazu/gms-backend/internal/module"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type auth struct {
	UserModule module.Auth
	Logger     zap.Logger
}

func Init(authModule module.Auth, log zap.Logger) handler.Auth {
	return &auth{
		UserModule: authModule,
		Logger:     log,
	}
}
func (u *auth) RegisterUserAuth(c *gin.Context) {
	var usr dto.UserAuth
	if err := c.ShouldBind(&usr); err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "unable to bind user to dto.User")
		_ = c.Error(err)
		return
	}
	err := u.UserModule.RegisterUserAuth(c, usr)
	if err != nil {
		_ = c.Error(err)
		return
	}
	response.SendSuccessResponse(c, http.StatusCreated, nil)

}
