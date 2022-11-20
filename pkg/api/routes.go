package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
Returned payload:
JSON string containing:

	  _id: string,
		name: string,
		icon: string,
		comment: string,
*/
func (app *App) getRewards(c *gin.Context) {

	cursor, err := Rewards.Find(context.TODO(), bson.D{})
	if err == mongo.ErrNoDocuments {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no rewards found"})
		return
	}
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		panic(err)
	}

	var rewards []Reward

	if err = cursor.All(context.TODO(), &rewards); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		panic(err)
	}

	jsonData, err := json.Marshal(rewards)
	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, string(jsonData))
}

/*
Expected payload:
JSON string containing:

	name: string,
	icon: string,
	comment: string,
*/
func (app *App) createReward(c *gin.Context) {
	var newReward NewReward

	// json.Unmarshal([]byte(rewardJson), &newReward)
	err := c.BindJSON(&newReward)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	doc := bson.D{
		{Key: "name", Value: newReward.Name},
		{Key: "icon", Value: newReward.Icon},
		{Key: "comment", Value: newReward.Comment},
	}

	result, err := Rewards.InsertOne(context.TODO(), doc)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		panic(err)
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		panic(err)
	}

	c.IndentedJSON(http.StatusCreated, string(jsonData))
}

/*
Expected payload:
JSON string containing:

	  id: string,
		name: string,
		icon: string,
		comment: string,
*/
func (app *App) updateReward(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Query("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		panic(err)
	}

	var updatedReward Reward

	// json.Unmarshal([]byte(rewardJson), &updatedReward)
	err = c.BindJSON(&updatedReward)
	if err != nil {
		return
	}

	// TODO Optimize by only updating changed fields
	update := bson.D{
		{
			"$set", bson.D{
				{Key: "name", Value: updatedReward.Name},
				{Key: "icon", Value: updatedReward.Icon},
				{Key: "comment", Value: updatedReward.Comment},
			},
		},
	}
	result, err := Rewards.UpdateOne(context.TODO(), bson.M{"_id": id}, update)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		panic(err)
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, string(jsonData))
}

/*
Expected payload:
String representing an ID
*/
func (app *App) deleteReward(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Query("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		panic(err)
	}

	result, err := Rewards.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		panic(err)
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, string(jsonData))
}
