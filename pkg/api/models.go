package api

import "go.mongodb.org/mongo-driver/bson/primitive"

type Reward struct {
	ID      primitive.ObjectID `bson:"_id"`
	Name    string
	Comment string
	Icon    string
}
