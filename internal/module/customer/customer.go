package customer

import (
	"context"
	"fmt"
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

type customer struct {
	log       *zap.Logger
	db        persistence.Customer
	userDB    persistence.User
	PaymentDB persistence.PaymentRule
}

func Init(db persistence.Customer, paymentdb persistence.PaymentRule, userdb persistence.User, log *zap.Logger) module.Customer {
	return &customer{
		log:       log,
		userDB:    userdb,
		db:        db,
		PaymentDB: paymentdb,
	}
}

func (u *customer) RegisterUser(ctx context.Context, customer dto.RegisterCustomer) (dto.RegisterCustomer, error) {
	validate := validator.New()
	if err := validate.Struct(&customer); err != nil {
		u.log.Error("validation error ", zap.Error(err))
		err = errors.ErrInvalidUserInput.Wrap(err, err.Error())
		return dto.RegisterCustomer{}, err
	}
	customer.ID = primitive.NewObjectID()
	customer.User.ID = primitive.NewObjectID()
	err := u.userDB.SaveUser(ctx, customer.User)
	if err != nil {
		return dto.RegisterCustomer{}, err
	}

	payment, err := u.PaymentDB.GetPaymentRule(ctx, bson.M{"_id": customer.PaymentRuleID})
	if err != nil {
		u.userDB.DeleteUser(ctx, customer.User.ID)
		return dto.RegisterCustomer{}, err
	}
	if len(payment) == 0 {
		u.userDB.DeleteUser(ctx, customer.User.ID)
		return dto.RegisterCustomer{}, fmt.Errorf("payment rule not found")
	}
	day := time.Hour * 24 * time.Duration(payment[0].NumberOfDays)
	endDate := customer.StarDate.Add(day)
	customer.EndDate = endDate
	customer.UserID = customer.User.ID
	customer.CreatedAt = time.Now()
	if err := u.db.SaveCustomer(ctx, customer); err != nil {
		u.userDB.DeleteUser(ctx, customer.User.ID)
		return dto.RegisterCustomer{}, err
	}
	return customer, nil
}

func (u *customer) GetCustomers(ctx context.Context, filter bson.M) ([]dto.RegisterCustomer, error) {
	id, ok := filter["id"]
	if !ok {
		return u.db.GetCustomer(ctx, filter)
	}
	hexID, _ := primitive.ObjectIDFromHex(id.(string))
	fil := bson.M{"_id": hexID}
	return u.db.GetCustomer(ctx, fil)

}

func (u *customer) UpdateCustomer(ctx context.Context, customerRequest dto.RegisterCustomer) (dto.RegisterCustomer, error) {

	userUpdate := bson.M{}
	if customerRequest.User.FirstName != "" {
		userUpdate["first_name"] = customerRequest.User.FirstName
	}
	if customerRequest.User.LastName != "" {
		userUpdate["last_name"] = customerRequest.User.FirstName

	}
	if customerRequest.User.Phone != "" {
		userUpdate["phone"] = customerRequest.User.Phone
	}
	if err := u.userDB.UpdateUser(ctx, bson.M{"_id": customerRequest.UserID}, userUpdate); err != nil {
		return dto.RegisterCustomer{}, err
	}

	if err := u.db.UpdateCustomer(ctx, bson.M{"_id": customerRequest.ID}, customerRequest); err != nil {
		return dto.RegisterCustomer{}, err
	}
	return customerRequest, nil
}

func (u *customer) DeleteCustomer(ctx context.Context, customerRequest dto.RegisterCustomer) error {
	return u.db.DeleteCustomer(ctx, customerRequest.ID)
}
