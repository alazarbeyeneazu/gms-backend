package customer

import (
	"context"

	"github.com/alazarbeyeneazu/gms-backend/internal/constants/errors"
	"github.com/alazarbeyeneazu/gms-backend/internal/constants/model/dto"
	"github.com/alazarbeyeneazu/gms-backend/internal/storage/persistence"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type customer struct {
	db  *mongo.Collection
	log zap.Logger
}

func Init(db *mongo.Collection, log zap.Logger) persistence.Customer {
	return &customer{
		db:  db,
		log: log,
	}
}
func (u *customer) SaveCustomer(ctx context.Context, customer dto.RegisterCustomer) error {
	_, err := u.db.InsertOne(ctx, customer)
	if err != nil {
		u.log.Error("Save customer", zap.Error(err))
		err = errors.ErrUnableTocreate.Wrap(err, err.Error(), "customer", customer)
		return err
	}
	return nil
}
func (u *customer) GetCustomer(ctx context.Context, filter bson.M) ([]dto.RegisterCustomer, error) {
	var customers []dto.RegisterCustomer
	result, err := u.db.Find(ctx, filter)
	if err != nil {
		u.log.Error("error while getting payment rule", zap.Error(err))
		err = errors.ErrUnableToUpdate.Wrap(err, err.Error(), zap.Any("filter", filter))
		return nil, err
	}

	defer result.Close(ctx)

	for result.Next(ctx) {
		var customer dto.RegisterCustomer
		if err := result.Decode(&customer); err != nil {
			u.log.Error("error decoding payment rule", zap.Error(err))
			err = errors.ErrUnableToUpdate.Wrap(err, err.Error(), zap.Any("filter", filter))
			return nil, err
		}
		customers = append(customers, customer)
	}

	if err := result.Err(); err != nil {
		u.log.Error("error while iterating payment rule result", zap.Error(err))
		err = errors.ErrUnableToUpdate.Wrap(err, err.Error(), zap.Any("filter", filter))
		return nil, err
	}

	return customers, nil
}
func (u *customer) UpdateCustomer(ctx context.Context, filter bson.M, update dto.RegisterCustomer) error {
	_, err := u.db.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		u.log.Error("error while updating customer", zap.Error(err))
		err = errors.ErrUnableToUpdate.Wrap(err, err.Error(), zap.Any("filter", filter), zap.Any("update", update))
		return err
	}
	return nil
}

func (u *customer) DeleteCustomer(ctx context.Context, customerID primitive.ObjectID) error {
	_, err := u.db.DeleteOne(ctx, bson.M{"_id": customerID})
	if err != nil {
		u.log.Error("error while updating customer", zap.Error(err))
		err = errors.ErrUnableToUpdate.Wrap(err, err.Error(), zap.Any("userID", customerID))
		return err
	}
	return nil
}
