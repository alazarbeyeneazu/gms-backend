package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id" `
	FirstName string             `json:"first_name" bson:"first_name" validate:"required"`
	LastName  string             `json:"last_name" bson:"last_name" validate:"required"`
	Phone     string             `json:"phone" bson:"phone" validate:"required"`
}
type Customer struct {
	UserID primitive.ObjectID `json:"user_id" bson:"user_id" validate:"required"`
}
type PaymentRule struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Type string             `json:"type" bson:"type"`
}
