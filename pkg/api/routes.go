package api

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getReward(name string) (reward *Reward, err error) {
	filter := bson.D{{"name", name}}

	var result Reward
	err = Rewards.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document found with the name %s\n", name)
		return nil, err
	}
	if err != nil {
		panic(err)
	}

	// jsonData, err := json.MarshalIndent(result, "", "    ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", jsonData)

	return &result, err
}

func addReward() {}

func updateReward() {}
