package users

import (
	"net/http"

	"github.com/alazarbeyeneazu/gms-backend/internal/constants/errors"
	"github.com/alazarbeyeneazu/gms-backend/internal/constants/model/dto"
	"github.com/alazarbeyeneazu/gms-backend/internal/constants/model/response"
	"github.com/alazarbeyeneazu/gms-backend/internal/handler"
	"github.com/alazarbeyeneazu/gms-backend/internal/module"
	"go.mongodb.org/mongo-driver/bson"

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
func (u *user) RegisterUser(c *gin.Context) {
	var usr dto.User
	if err := c.ShouldBind(&usr); err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "unable to bind user to dto.User")
		_ = c.Error(err)
		return
	}
	res, err := u.UserModule.RegisterUser(c, usr)
	if err != nil {
		_ = c.Error(err)
		return
	}
	response.SendSuccessResponse(c, http.StatusCreated, res)

}
func (u *user) UpdateUser(c *gin.Context) {
	var usr dto.User
	if err := c.ShouldBind(&usr); err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "unable to bind user to dto.User")
		_ = c.Error(err)
		return
	}
	res, err := u.UserModule.UpdateUser(c, usr)
	if err != nil {
		_ = c.Error(err)
		return
	}
	response.SendSuccessResponse(c, http.StatusOK, res)

}
func (u *user) DeleteUser(c *gin.Context) {
	var usr dto.User
	if err := c.ShouldBind(&usr); err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "unable to bind user to dto.User")
		_ = c.Error(err)
		return
	}
	res, err := u.UserModule.DeleteUser(c, usr)
	if err != nil {
		_ = c.Error(err)
		return
	}
	response.SendSuccessResponse(c, http.StatusOK, res)

}

func (u *user) GetUsers(c *gin.Context) {
	var bsonMap bson.M
	filter := c.Query("filter")
	err := bson.UnmarshalExtJSON([]byte(filter), true, &bsonMap)
	if err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "unable to bind user to dto.user")
		_ = c.Error(err)
		return
	}
	res, err := u.UserModule.GetUsers(c, bsonMap)
	if err != nil {
		_ = c.Error(err)
		return
	}
	response.SendSuccessResponse(c, http.StatusOK, res)

}
