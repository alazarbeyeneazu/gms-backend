package paymentrule

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

type paymentRule struct {
	db  *mongo.Collection
	log zap.Logger
}

func Init(db *mongo.Collection, log zap.Logger) persistence.PaymentRule {
	return &paymentRule{
		db:  db,
		log: log,
	}
}
func (u *paymentRule) SavepaymentRule(ctx context.Context, paymentRule dto.PaymentRule) error {
	_, err := u.db.InsertOne(ctx, paymentRule)
	if err != nil {
		u.log.Error("Save payment rule", zap.Error(err))
		err = errors.ErrUnableTocreate.Wrap(err, err.Error(), "payment rule", paymentRule)
		return err
	}
	return nil
}
func (u *paymentRule) UpdatepaymentRule(ctx context.Context, filter bson.M, update bson.M) error {
	_, err := u.db.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		u.log.Error("error while updating payment rule", zap.Error(err))
		err = errors.ErrUnableToUpdate.Wrap(err, err.Error(), zap.Any("filter", filter), zap.Any("update", update))
		return err
	}
	return nil
}
func (u *paymentRule) GetPaymentRule(ctx context.Context, filter bson.M) ([]dto.PaymentRule, error) {
	var paymentRules []dto.PaymentRule
	result, err := u.db.Find(ctx, filter)
	if err != nil {
		u.log.Error("error while getting payment rule", zap.Error(err))
		err = errors.ErrUnableToUpdate.Wrap(err, err.Error(), zap.Any("filter", filter))
		return nil, err
	}

	defer result.Close(ctx)

	for result.Next(ctx) {
		var paymentRule dto.PaymentRule
		if err := result.Decode(&paymentRule); err != nil {
			u.log.Error("error decoding payment rule", zap.Error(err))
			err = errors.ErrUnableToUpdate.Wrap(err, err.Error(), zap.Any("filter", filter))
			return nil, err
		}
		paymentRules = append(paymentRules, paymentRule)
	}

	if err := result.Err(); err != nil {
		u.log.Error("error while iterating payment rule result", zap.Error(err))
		err = errors.ErrUnableToUpdate.Wrap(err, err.Error(), zap.Any("filter", filter))
		return nil, err
	}

	return paymentRules, nil
}
func (u *paymentRule) DeletepaymentRule(ctx context.Context, paymentID primitive.ObjectID) error {
	_, err := u.db.DeleteOne(ctx, bson.M{"_id": paymentID})
	if err != nil {
		u.log.Error("error while updating user", zap.Error(err))
		err = errors.ErrUnableToUpdate.Wrap(err, err.Error(), zap.Any("paymentID", paymentID))
		return err
	}
	return nil
}
