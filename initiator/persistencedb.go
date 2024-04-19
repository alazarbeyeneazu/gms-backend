package initiator

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Persistence struct {
}

func InitPersistence(db *mongo.Client, log zap.Logger) Persistence {
	return Persistence{}
}
