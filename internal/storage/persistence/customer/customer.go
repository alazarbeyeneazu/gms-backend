package customer

import (
	"context"

	"github.com/alazarbeyeneazu/gms-backend/internal/constants/errors"
	"github.com/alazarbeyeneazu/gms-backend/internal/constants/model/dto"
	"github.com/alazarbeyeneazu/gms-backend/internal/storage/persistence"
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
func (u *customer) SaveCustomer(ctx context.Context, user dto.User) error {
	_, err := u.db.InsertOne(ctx, user)
	if err != nil {
		u.log.Error("Save User", zap.Error(err))
		err = errors.ErrUnableTocreate.Wrap(err, err.Error(), "user", user)
		return err
	}
	return nil
}
