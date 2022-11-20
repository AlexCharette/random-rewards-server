package api

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Rewards *mongo.Collection
)

func (app *App) Initialize() {
	uri := os.Getenv("DB_URI")
	if uri == "" {
		log.Fatal("No environment variable found named DB_URI")
	}
	err := error(nil)

	monitor := &event.CommandMonitor{
		Started: func(_ context.Context, e *event.CommandStartedEvent) {
			fmt.Println(e.Command)
		},
		Succeeded: func(_ context.Context, e *event.CommandSucceededEvent) {
			fmt.Println(e.Reply)
		},
		Failed: func(_ context.Context, e *event.CommandFailedEvent) {
			fmt.Println(e.Failure)
		},
	}

	opts := options.Client().SetMonitor(monitor).ApplyURI(uri)

	app.DBClient, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	Rewards = app.DBClient.Database(os.Getenv("DB_NAME")).Collection("rewards")

	app.initializeRoutes()
}

func (app *App) Run() {

	app.Router.Run("localhost:8080")

	defer func() {
		if err := app.DBClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}

func (app *App) initializeRoutes() {
	app.Router = gin.Default()
	app.Router.GET("/rewards", app.getRewards)
	app.Router.POST("/reward", app.createReward)
	app.Router.PUT("/reward", app.updateReward)
	app.Router.DELETE("/reward", app.deleteReward)
}
