package auth

import (
	"context"
	"fmt"

	"github.com/alazarbeyeneazu/gms-backend/internal/constants/errors"
	"github.com/alazarbeyeneazu/gms-backend/internal/constants/model/dto"
	"github.com/alazarbeyeneazu/gms-backend/internal/module"
	"github.com/alazarbeyeneazu/gms-backend/internal/storage/persistence"
	"github.com/alazarbeyeneazu/gms-backend/platform/utils"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type userAuth struct {
	log             *zap.Logger
	userDB          persistence.User
	authPersistence persistence.Auth
}

func Init(userdb persistence.User, authdb persistence.Auth, log *zap.Logger) module.Auth {
	return &userAuth{
		log:             log,
		userDB:          userdb,
		authPersistence: authdb,
	}
}
func (u *userAuth) RegisterUserAuth(ctx context.Context, userAuth dto.UserAuth) error {
	validate := validator.New()
	if err := validate.Struct(&userAuth); err != nil {
		u.log.Error("validation error ", zap.Error(err))
		err = errors.ErrInvalidUserInput.Wrap(err, err.Error())
		return err
	}
	user, err := u.userDB.GetUsers(ctx, bson.M{"_id": userAuth.UserID})
	if err != nil {
		return err
	}
	if len(user) == 0 {
		err = fmt.Errorf("user not found")
		u.log.Error("validation error ", zap.Error(err))
		err = errors.ErrInvalidUserInput.Wrap(err, err.Error())
		return err
	}
	pass, err := utils.GenerateHash(userAuth.Password)
	if err != nil {
		err = fmt.Errorf("unable to hash user password")
		u.log.Error("validation error ", zap.Error(err))
		err = errors.ErrUnExpectedError.Wrap(err, err.Error())
		return err
	}
	authp := dto.UserAuth{
		ID:       primitive.NewObjectID(),
		UserID:   user[0].ID,
		Password: pass,
	}
	return u.authPersistence.SaveUserAuth(ctx, authp)

}
