package dto

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
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

type RegisterCustomer struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	User          User               `json:"user" validate:"required"`
	UserID        primitive.ObjectID `json:"user_id" bson:"user_id"`
	PaymentRuleID primitive.ObjectID `json:"payment_rule_id" bson:"payment_rule_id" validate:"required"`
	QRCode        string             `json:"qr_code" bson:"qr_code" validate:"required"`
	StarDate      time.Time          `json:"start_date" bson:"start_date" validate:"required"`
	EndDate       time.Time          `json:"end_date" bson:"end_date"`
	RegisterdBy   primitive.ObjectID `json:"registered_by" bson:"registered_by"`
	CreatedAt     time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at"`
	DeletedAt     time.Time          `json:"deleted_at" bson:"deleted_at"`
}
type UserAuth struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id" validate:"required"`
	Password string             `json:"password" bson:"password" validate:"required"`
}
type Claim struct {
	UserID primitive.ObjectID `json:"user_id"`
	jwt.RegisteredClaims
}
type LoginRequest struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}
