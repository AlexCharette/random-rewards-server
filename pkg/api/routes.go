package api

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// TODO Define expected inputs and outputs

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

func addReward(rewardJson string) (reward *Reward, err error) {
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

	return &result, err
}

func updateReward(rewardJason string) {
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
}

func deleteReward(id string) {
	filter := bson.D{
		{"_id", id},
	}
	result, err := Rewards.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
}
