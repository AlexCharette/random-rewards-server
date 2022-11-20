package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	api "random-rewards/pkg/api"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var dummyRewards = []api.NewReward{
	{Name: "Bramibam's Donut", Icon: "🍩", Comment: "You get a delicious donut!"},
	{Name: "Barebells Bar", Icon: "🍫", Comment: "Get stronger with protein bar"},
	{Name: "Doner Kebab", Icon: "🥙", Comment: "Treat yourself to a kebab"},
}

var app = api.App{}

func setupSuite(tb testing.TB) func(tb testing.TB) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	dbName := os.Getenv("DB_NAME")
	collName := "rewards"

	app.Initialize()
	checkDatabase(app.DBClient, dbName)
	checkCollection(app.DBClient.Database(dbName), collName)
	app.Run()

	// Return a function to teardown the test
	return func(tb testing.TB) {
		clearCollection(app.DBClient.Database(dbName), collName)
	}
}

func checkDatabase(client *mongo.Client, name string) {
	dbs, err := client.ListDatabaseNames(context.TODO(), bson.D{
		{Key: "name", Value: name},
	})
	if err != nil {
		panic(err)
	}

	// If the requested database doesn't exist
	if len(dbs) == 0 {
		log.Fatal(fmt.Printf("No database found by the name %s", name))
	}
}

func checkCollection(database *mongo.Database, name string) {
	colls, err := database.ListCollectionNames(context.TODO(), bson.D{
		{Key: "name", Value: name},
	})
	if err != nil {
		panic(err)
	}

	// If the requested collection doesn't exist
	if len(colls) == 0 {
		database.CreateCollection(context.TODO(), name)
	}
}

func clearCollection(database *mongo.Database, name string) {
	err := database.Collection(name).Drop(context.TODO())
	if err != nil {
		panic(err)
	}
}

func TestGetRewards(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	log.Println("No .env file found")
	req, _ := http.NewRequest("GET", "/rewards", nil)
	w := httptest.NewRecorder()
	app.Router.ServeHTTP(w, req)

	var rewards []api.Reward
	json.Unmarshal(w.Body.Bytes(), &rewards)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, rewards)
}
