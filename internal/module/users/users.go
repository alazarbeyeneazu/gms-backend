package users

import (
	"context"

	"github.com/alazarbeyeneazu/gms-backend/internal/constants/errors"
	"github.com/alazarbeyeneazu/gms-backend/internal/constants/model/dto"
	"github.com/alazarbeyeneazu/gms-backend/internal/module"
	"github.com/alazarbeyeneazu/gms-backend/internal/storage/persistence"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type user struct {
	log    *zap.Logger
	userDB persistence.User
}

func Init(userdb persistence.User, log *zap.Logger) module.User {
	return &user{
		log:    log,
		userDB: userdb,
	}
}
func (u *user) RegisterUser(ctx context.Context, user dto.User) (dto.User, error) {
	validate := validator.New()
	if err := validate.Struct(&user); err != nil {
		u.log.Error("validation error ", zap.Error(err))
		err = errors.ErrInvalidUserInput.Wrap(err, err.Error())
		return dto.User{}, err
	}
	user.ID = primitive.NewObjectID()
	if err := u.userDB.SaveUser(ctx, user); err != nil {
		return dto.User{}, err
	}
	return user, nil
}
func (u *user) UpdateUser(ctx context.Context, user dto.User) (dto.User, error) {

	update := bson.M{}
	if user.FirstName != "" {
		update["first_name"] = user.FirstName
	}
	if user.LastName != "" {
		update["last_name"] = user.LastName
	}
	if user.Phone != "" {
		update["phone"] = user.Phone
	}

	u.userDB.UpdateUser(ctx, bson.M{"_id": user.ID}, update)
	return user, nil
}
func (u *user) DeleteUser(ctx context.Context, user dto.User) (dto.User, error) {
	u.userDB.DeleteUser(ctx, user.ID)
	return user, nil
}
func (u *user) GetUsers(ctx context.Context, filter bson.M) ([]dto.User, error) {
	return u.userDB.GetUsers(ctx, filter)
}
