package users

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

type user struct {
	db  *mongo.Collection
	log zap.Logger
}

func Init(db *mongo.Collection, log zap.Logger) persistence.User {
	return &user{
		db:  db,
		log: log,
	}
}
func (u *user) SaveUser(ctx context.Context, user dto.User) error {
	_, err := u.db.InsertOne(ctx, user)
	if err != nil {
		u.log.Error("Save User", zap.Error(err))
		err = errors.ErrUnableTocreate.Wrap(err, err.Error(), "user", user)
		return err
	}
	return nil
}
func (u *user) UpdateUser(ctx context.Context, filter bson.M, update bson.M) error {
	_, err := u.db.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		u.log.Error("error while updating user", zap.Error(err))
		err = errors.ErrUnableToUpdate.Wrap(err, err.Error(), zap.Any("filter", filter), zap.Any("update", update))
		return err
	}
	return nil
}
func (u *user) DeleteUser(ctx context.Context, userID primitive.ObjectID) error {
	_, err := u.db.DeleteOne(ctx, bson.M{"_id": userID})
	if err != nil {
		u.log.Error("error while updating user", zap.Error(err))
		err = errors.ErrUnableToUpdate.Wrap(err, err.Error(), zap.Any("userID", userID))
		return err
	}
	return nil
}
func (u *user) GetUsers(ctx context.Context, filter bson.M) ([]dto.User, error) {
	var users []dto.User
	result, err := u.db.Find(ctx, filter)
	if err != nil {
		u.log.Error("error while getting payment rule", zap.Error(err))
		err = errors.ErrUnableToUpdate.Wrap(err, err.Error(), zap.Any("filter", filter))
		return nil, err
	}

	defer result.Close(ctx)

	for result.Next(ctx) {
		var usr dto.User
		if err := result.Decode(&usr); err != nil {
			u.log.Error("error decoding payment rule", zap.Error(err))
			err = errors.ErrUnableToUpdate.Wrap(err, err.Error(), zap.Any("filter", filter))
			return nil, err
		}
		users = append(users, usr)
	}

	if err := result.Err(); err != nil {
		u.log.Error("error while iterating payment rule result", zap.Error(err))
		err = errors.ErrUnableToUpdate.Wrap(err, err.Error(), zap.Any("filter", filter))
		return nil, err
	}

	return users, nil
}
