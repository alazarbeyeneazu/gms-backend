package customer

import (
	"github.com/alazarbeyeneazu/gms-backend/internal/storage/persistence"
	"go.uber.org/zap"
)

type customer struct {
	log    *zap.Logger
	userDB persistence.User
}

// func Init(userdb persistence.User, log *zap.Logger) module.Customer {
// 	return &customer{
// 		log:    log,
// 		userDB: userdb,
// 	}
// }
// func (u *customer) RegisterUser(ctx context.Context, user dto.User) (dto.User, error) {
// 	validate := validator.New()
// 	if err := validate.Struct(&user); err != nil {
// 		u.log.Error("validation error ", zap.Error(err))
// 		err = errors.ErrInvalidUserInput.Wrap(err, err.Error())
// 		return dto.User{}, err
// 	}
// 	user.ID = primitive.NewObjectID()
// 	if err := u.userDB.SaveUser(ctx, user); err != nil {
// 		return dto.User{}, err
// 	}
// 	return user, nil
// }
