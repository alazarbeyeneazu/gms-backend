package module

import (
	"context"

	"github.com/alazarbeyeneazu/gms-backend/internal/constants/model/dto"
)

type User interface {
	RegisterUser(ctxt context.Context, user dto.User) (dto.User, error)
	UpdateUser(ctx context.Context, user dto.User) (dto.User, error)
	DeleteUser(ctx context.Context, user dto.User) (dto.User, error)
}
type Customer interface {
}
