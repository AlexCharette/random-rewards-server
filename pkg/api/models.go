package api

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	DBClient *mongo.Client
	Router   *gin.Engine
}

type Reward struct {
	ID      primitive.ObjectID `bson:"_id"`
	Name    string             `bson:"name"`
	Comment string             `bson:"comment"`
	Icon    string             `bson:"icon"`
}

type NewReward struct {
	Name    string `bson:"name"`
	Comment string `bson:"comment"`
	Icon    string `bson:"icon"`
}
