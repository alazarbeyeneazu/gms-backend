package persistence

import (
	"context"

	"github.com/alazarbeyeneazu/gms-backend/internal/constants/model/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User interface {
	SaveUser(ctx context.Context, user dto.User) error
	UpdateUser(ctx context.Context, filter bson.M, update bson.M) error
	DeleteUser(ctx context.Context, userID primitive.ObjectID) error
	GetUsers(ctx context.Context, filter bson.M) ([]dto.User, error)
}
type Customer interface {
}
type PaymentRule interface {
	SavepaymentRule(ctx context.Context, paymentRule dto.PaymentRule) error
	UpdatepaymentRule(ctx context.Context, filter bson.M, update bson.M) error
	GetPaymentRule(ctx context.Context, filter bson.M) ([]dto.PaymentRule, error)
	DeletepaymentRule(ctx context.Context, paymentID primitive.ObjectID) error
}
