package initiator

import (
	"github.com/alazarbeyeneazu/gms-backend/internal/constants/model/persistancedb"
	"github.com/alazarbeyeneazu/gms-backend/internal/storage/persistence"
	"github.com/alazarbeyeneazu/gms-backend/internal/storage/persistence/users"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Persistence struct {
	user persistence.User
}

func InitPersistence(db *mongo.Client, log zap.Logger) Persistence {
	return Persistence{
		user: users.Init(persistancedb.GetCollection(db, "users"), log),
	}
}
