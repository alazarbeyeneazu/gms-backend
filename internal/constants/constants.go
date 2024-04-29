package constants

import "go.mongodb.org/mongo-driver/bson/primitive"

type Auth struct {
	UserID primitive.ObjectID `json:"user_id"`
}
