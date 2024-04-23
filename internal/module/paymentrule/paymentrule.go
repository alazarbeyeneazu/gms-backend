package paymentrule

import (
	"context"
	"time"

	"github.com/alazarbeyeneazu/gms-backend/internal/constants/errors"
	"github.com/alazarbeyeneazu/gms-backend/internal/constants/model/dto"
	"github.com/alazarbeyeneazu/gms-backend/internal/module"
	"github.com/alazarbeyeneazu/gms-backend/internal/storage/persistence"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type paymentRule struct {
	log           *zap.Logger
	paymentRuleDB persistence.PaymentRule
}

func Init(paymentRuledb persistence.PaymentRule, log *zap.Logger) module.PaymentRule {
	return &paymentRule{
		log:           log,
		paymentRuleDB: paymentRuledb,
	}
}
func (u *paymentRule) CreatePaymentRule(ctx context.Context, paymentRule dto.PaymentRule) (dto.PaymentRule, error) {
	validate := validator.New()
	if err := validate.Struct(&paymentRule); err != nil {
		u.log.Error("validation error ", zap.Error(err))
		err = errors.ErrInvalidUserInput.Wrap(err, err.Error())
		return dto.PaymentRule{}, err
	}

	paymentRule.ID = primitive.NewObjectID()
	paymentRule.CreatedAt = time.Now()
	if err := u.paymentRuleDB.SavepaymentRule(ctx, paymentRule); err != nil {
		return dto.PaymentRule{}, err
	}
	return paymentRule, nil
}

func (u *paymentRule) UpdatepaymentRule(ctx context.Context, paymentRule dto.PaymentRule) (dto.PaymentRule, error) {

	update := bson.M{}
	if paymentRule.Type != "" {
		update["type"] = paymentRule.Type
	}
	if paymentRule.Payment > 0 {
		update["payment"] = paymentRule.Payment
	}
	if paymentRule.DaysPerWeek != 0 {
		update["days_per_week"] = paymentRule.DaysPerWeek
	}
	if paymentRule.TimesPerDay != 0 {
		update["times_per_day"] = paymentRule.TimesPerDay
	}
	if paymentRule.NumberOfDays != 0 {
		update["number_of_days"] = paymentRule.NumberOfDays
	}
	paymentRule.UpdatedAt = time.Now()
	u.paymentRuleDB.UpdatepaymentRule(ctx, bson.M{"_id": paymentRule.ID}, update)
	return paymentRule, nil
}
func (u *paymentRule) GetPaymentRule(ctx context.Context, filter bson.M) ([]dto.PaymentRule, error) {
	return u.paymentRuleDB.GetPaymentRule(ctx, filter)
}

func (u *paymentRule) DeletepaymentRule(ctx context.Context, paymentRule dto.PaymentRule) (dto.PaymentRule, error) {
	u.paymentRuleDB.DeletepaymentRule(ctx, paymentRule.ID)
	return paymentRule, nil
}
