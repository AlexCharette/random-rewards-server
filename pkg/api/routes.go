package api

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// TODO Define expected inputs and outputs

func getReward(name string) (resultJson string, err error) {

	// TODO Take JSON from param

	filter := bson.D{{"name", name}}

	var result Reward
	err = Rewards.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document found with the name %s\n", name)
		return "", err
	}
	if err != nil {
		panic(err)
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	return string(jsonData), err
}

func addReward(rewardJson string) (resultJson string, err error) {
	var newReward Reward
	// TODO Handle errors
	json.Unmarshal([]byte(rewardJson), &newReward)
	doc := bson.D{
		{"name", newReward.Name},
		{"comment", newReward.Comment},
		{"icon", newReward.Icon},
	}

	result, err := Rewards.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	return string(jsonData), err
}

func updateReward(rewardJason string) (resultJson string, err error) {
	var updatedReward Reward
	json.Unmarshal([]byte(rewardJason), &updatedReward)
	filter := bson.D{
		{"_id", updatedReward.ID},
	}
	// TODO Optimize by only updating changed fields
	update := bson.D{
		{
			"$set", bson.D{
				{"name", updatedReward.Name},
				{"comment", updatedReward.Comment},
				{"icon", updatedReward.Icon},
			},
		},
	}
	result, err := Rewards.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	return string(jsonData), err
}

func deleteReward(id string) (resultJson string, err error) {
	filter := bson.D{
		{"_id", id},
	}
	result, err := Rewards.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	return string(jsonData), err
}
