package auth

import (
	"context"

	"github.com/alazarbeyeneazu/gms-backend/internal/constants/errors"
	"github.com/alazarbeyeneazu/gms-backend/internal/constants/model/dto"
	"github.com/alazarbeyeneazu/gms-backend/internal/storage/persistence"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type auth struct {
	db  *mongo.Collection
	log zap.Logger
}

func Init(db *mongo.Collection, log zap.Logger) persistence.Auth {
	return &auth{
		db:  db,
		log: log,
	}
}
func (u *auth) SaveUserAuth(ctx context.Context, userAuth dto.UserAuth) error {
	_, err := u.db.InsertOne(ctx, userAuth)
	if err != nil {
		u.log.Error("Save User", zap.Error(err))
		err = errors.ErrUnableTocreate.Wrap(err, err.Error(), "user auth", userAuth)
		return err
	}
	return nil
}
func (u *auth) GetAuth(ctx context.Context, filter bson.M) (dto.UserAuth, error) {
	var authdata dto.UserAuth
	err := u.db.FindOne(ctx, filter).Decode(&authdata)
	if err != nil {
		u.log.Error("Save User", zap.Error(err))
		err = errors.ErrUnableToGet.Wrap(err, err.Error(), "user auth", authdata)
		return dto.UserAuth{}, err
	}

	return authdata, nil
}
