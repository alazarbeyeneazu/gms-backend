package module

import (
	"context"

	"github.com/alazarbeyeneazu/gms-backend/internal/constants/model/dto"
	"go.mongodb.org/mongo-driver/bson"
)

type User interface {
	RegisterUser(ctxt context.Context, user dto.User) (dto.User, error)
	UpdateUser(ctx context.Context, user dto.User) (dto.User, error)
	DeleteUser(ctx context.Context, user dto.User) (dto.User, error)
	GetUsers(ctx context.Context, filter bson.M) ([]dto.User, error)
}
type PaymentRule interface {
	CreatePaymentRule(ctx context.Context, paymentRule dto.PaymentRule) (dto.PaymentRule, error)
	UpdatepaymentRule(ctx context.Context, paymentRule dto.PaymentRule) (dto.PaymentRule, error)
	GetPaymentRule(ctx context.Context, filter bson.M) ([]dto.PaymentRule, error)
	DeletepaymentRule(ctx context.Context, paymentRule dto.PaymentRule) (dto.PaymentRule, error)
}
