package paymentrule

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

type paymentRule struct {
	paymentRuleModule module.PaymentRule
	Logger            zap.Logger
}

func Init(paymentRuleModule module.PaymentRule, log zap.Logger) handler.PaymentRule {
	return &paymentRule{
		paymentRuleModule: paymentRuleModule,
		Logger:            log,
	}
}
func (u *paymentRule) CreatePaymentRule(c *gin.Context) {
	var payment dto.PaymentRule
	if err := c.ShouldBind(&payment); err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "unable to bind user to dto.PaymentRule")
		_ = c.Error(err)
		return
	}
	res, err := u.paymentRuleModule.CreatePaymentRule(c, payment)
	if err != nil {
		_ = c.Error(err)
		return
	}
	response.SendSuccessResponse(c, http.StatusCreated, res)

}
func (u *paymentRule) UpdatepaymentRule(c *gin.Context) {
	var payment dto.PaymentRule
	if err := c.ShouldBind(&payment); err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "unable to bind user to dto.PaymentRule")
		_ = c.Error(err)
		return
	}
	res, err := u.paymentRuleModule.UpdatepaymentRule(c, payment)
	if err != nil {
		_ = c.Error(err)
		return
	}
	response.SendSuccessResponse(c, http.StatusOK, res)

}
func (u *paymentRule) GetPaymentRule(c *gin.Context) {
	var bsonMap bson.M
	filter := c.Query("filter")
	err := bson.UnmarshalExtJSON([]byte(filter), true, &bsonMap)
	if err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "unable to bind user to dto.PaymentRule")
		_ = c.Error(err)
		return
	}
	res, err := u.paymentRuleModule.GetPaymentRule(c, bsonMap)
	if err != nil {
		_ = c.Error(err)
		return
	}
	response.SendSuccessResponse(c, http.StatusOK, res)

}

func (u *paymentRule) DeletepaymentRule(c *gin.Context) {
	var payment dto.PaymentRule
	if err := c.ShouldBind(&payment); err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "unable to bind user to dto.PaymentRule")
		_ = c.Error(err)
		return
	}
	res, err := u.paymentRuleModule.DeletepaymentRule(c, payment)
	if err != nil {
		_ = c.Error(err)
		return
	}
	response.SendSuccessResponse(c, http.StatusOK, res)

}
