package initiator

import (
	"github.com/alazarbeyeneazu/gms-backend/internal/constants/model/persistancedb"
	"github.com/alazarbeyeneazu/gms-backend/internal/storage/persistence"
	"github.com/alazarbeyeneazu/gms-backend/internal/storage/persistence/auth"
	"github.com/alazarbeyeneazu/gms-backend/internal/storage/persistence/customer"
	"github.com/alazarbeyeneazu/gms-backend/internal/storage/persistence/paymentrule"
	"github.com/alazarbeyeneazu/gms-backend/internal/storage/persistence/users"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Persistence struct {
	user        persistence.User
	paymentRule persistence.PaymentRule
	customer    persistence.Customer
	auth        persistence.Auth
}

func InitPersistence(db *mongo.Client, log zap.Logger) Persistence {
	return Persistence{
		user:        users.Init(persistancedb.GetCollection(db, "users"), log),
		paymentRule: paymentrule.Init(persistancedb.GetCollection(db, "payment_rule"), log),
		customer:    customer.Init(persistancedb.GetCollection(db, "customer"), log),
		auth:        auth.Init(persistancedb.GetCollection(db, "auth"), log),
	}
}
