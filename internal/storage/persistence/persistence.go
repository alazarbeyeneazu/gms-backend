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
}
type Customer interface {
}
type PaymentRule interface{}
