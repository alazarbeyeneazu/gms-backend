package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/alazarbeyeneazu/gms-backend/internal/constants/errors"
	"github.com/alazarbeyeneazu/gms-backend/internal/constants/model/dto"
	"github.com/alazarbeyeneazu/gms-backend/internal/module"
	"github.com/alazarbeyeneazu/gms-backend/internal/storage/persistence"
	"github.com/alazarbeyeneazu/gms-backend/platform/utils"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type userAuth struct {
	log             *zap.Logger
	userDB          persistence.User
	authPersistence persistence.Auth
	authKey         string
}

func Init(userdb persistence.User, authdb persistence.Auth, authkey string, log *zap.Logger) module.Auth {
	return &userAuth{
		log:             log,
		userDB:          userdb,
		authPersistence: authdb,
		authKey:         authkey,
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
func (u *userAuth) Login(ctx context.Context, loginRequst dto.LoginRequest) (string, error) {
	validate := validator.New()
	if err := validate.Struct(&loginRequst); err != nil {
		u.log.Error("validation error ", zap.Error(err))
		err = errors.ErrInvalidUserInput.Wrap(err, err.Error())
		return "", err
	}
	user, err := u.userDB.GetUsers(ctx, bson.M{"phone": loginRequst.Phone})
	if err != nil {
		u.log.Error("incorrect username or password", zap.Error(err))
		err = errors.ErrInvalidUserInput.Wrap(err, "(የተሳሳተ ስልክ ወይም የይለፍ ቃል) incorrect username or password")
		return "", err
	}
	password, err := u.authPersistence.GetAuth(ctx, bson.M{"user_id": user[0].ID})
	if err != nil {
		return "", err
	}
	if err := utils.ComparePasswords(password.Password, loginRequst.Password); err != nil {
		u.log.Error("(የተሳሳተ ስልክ ወይም የይለፍ ቃል) incorrect phone or password ", zap.Error(err))
		err = errors.ErrInvalidUserInput.Wrap(err, "(የተሳሳተ ስልክ ወይም የይለፍ ቃል) incorrect username or password")
		return "", err
	}
	expirationTime := time.Now().Add(time.Hour * 16)
	claims := &dto.Claim{
		UserID: password.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(u.authKey))
	if err != nil {
		u.log.Error("validation error ", zap.Error(err))
		err = errors.ErrUnExpectedError.Wrap(err, err.Error())
		return "", err
	}
	return tokenString, nil

}
