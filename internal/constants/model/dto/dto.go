package dto

import (
	"time"

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
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Type         string             `json:"type" bson:"type" validate:"required"`
	Payment      float64            `json:"payment" bson:"payment" validate:"required"`
	DaysPerWeek  int                `json:"days_per_week" bson:"days_per_week" validate:"required"`
	TimesPerDay  int                `json:"times_per_day" bson:"times_per_day" validate:"required"`
	NumberOfDays int                `json:"number_of_days" bson:"number_of_days" validate:"required"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
	DeletedAt    time.Time          `json:"deleted_at" bson:"deleted_at"`
}
