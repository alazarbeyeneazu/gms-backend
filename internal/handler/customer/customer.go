package customer

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

type customer struct {
	customerModule module.Customer
	Logger         zap.Logger
}

func Init(customerModule module.Customer, log zap.Logger) handler.Customer {
	return &customer{
		customerModule: customerModule,
		Logger:         log,
	}
}
func (u *customer) RegisterCustomer(c *gin.Context) {
	var customer dto.RegisterCustomer
	if err := c.ShouldBind(&customer); err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "unable to bind user to dto.PaymentRule")
		_ = c.Error(err)
		return
	}
	res, err := u.customerModule.RegisterUser(c, customer)
	if err != nil {
		_ = c.Error(err)
		return
	}
	response.SendSuccessResponse(c, http.StatusCreated, res)

}

func (u *customer) UpdateCustomer(c *gin.Context) {
	var payment dto.RegisterCustomer
	if err := c.ShouldBind(&payment); err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "unable to bind user to dto.PaymentRule")
		_ = c.Error(err)
		return
	}
	res, err := u.customerModule.UpdateCustomer(c, payment)
	if err != nil {
		_ = c.Error(err)
		return
	}
	response.SendSuccessResponse(c, http.StatusOK, res)

}
func (u *customer) GetCustomer(c *gin.Context) {
	var bsonMap bson.M
	filter := c.Query("filter")
	err := bson.UnmarshalExtJSON([]byte(filter), true, &bsonMap)
	if err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "unable to bind user to dto.Customer")
		_ = c.Error(err)
		return
	}
	res, err := u.customerModule.GetCustomers(c, bsonMap)
	if err != nil {
		_ = c.Error(err)
		return
	}
	response.SendSuccessResponse(c, http.StatusOK, res)

}

func (u *customer) DeleteCustomer(c *gin.Context) {
	var payment dto.RegisterCustomer
	if err := c.ShouldBind(&payment); err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "unable to bind user to dto.RegisterCustomer")
		_ = c.Error(err)
		return
	}
	err := u.customerModule.DeleteCustomer(c, payment)
	if err != nil {
		_ = c.Error(err)
		return
	}
	response.SendSuccessResponse(c, http.StatusOK, nil)

}
