package main

import (
	"context"
	"net/http"
	"time"
	"vaas/structs"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func addUser(c *gin.Context) {
	userId := c.Param("id")
	userName := c.Param("username")
	user := structs.User{UserID: userId, UserName: userName, Games: []string{}, NumGames: 0, TotalGuesses: 0}

	database := client.Database("VaasDatabase")
	userCollection := database.Collection("users")

	var existingUser structs.User
	criteria := map[string]interface{}{
		"username": userName,
		"userid":   userId,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := userCollection.FindOne(ctx, criteria).Decode(&existingUser)

	if err == mongo.ErrNoDocuments {
		userCollection.InsertOne(context.TODO(), user)

		c.JSON(http.StatusOK, map[string]string{"message": "User " + userName + " created"})
	} else {
		c.JSON(http.StatusOK, map[string]string{"message": "User " + userName + " already exists"})
	}
}

func getUser(c *gin.Context) {
	userName := c.Param("username")

	database := client.Database("VaasDatabase")
	userCollection := database.Collection("users")

	matchStage := bson.A{
		bson.D{{"$match", bson.D{{"username", userName}}}},
	}
	cursor, err := userCollection.Aggregate(context.TODO(), matchStage)
	defer cursor.Close(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve game from mongo: " + err.Error()})
		return
	}
	var user structs.User
	if cursor.Next(context.TODO()) {
		if err := cursor.Decode(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decode user from mongo: " + err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, user)
}
