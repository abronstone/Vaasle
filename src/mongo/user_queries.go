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

func newUser(c *gin.Context) {
	/*
		Adds a new user to the 'users' collection in the database.
		Expects a JSON payload with userName and id fields.

		@return: confirmation message
	*/

	// Bind incoming JSON to UserRequestBody struct
	var requestBody structs.NewUserRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Populate the User struct with received userName and ID
	var user structs.User
	user.UserName = requestBody.UserName
	user.Id = requestBody.Id

	// Connect to MongoDB (replace client with your actual client)
	database := client.Database("VaasDatabase")
	userCollection := database.Collection("users")

	// Check if the user already exists
	cursor, err := userCollection.Find(context.TODO(), bson.M{"id": user.Id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.TODO())

	// If no existing user is found, insert new user
	if !cursor.Next(context.TODO()) {
		_, err := userCollection.InsertOne(context.TODO(), user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, structs.Message{Message: "User inserted into database"})
	} else {
		c.JSON(http.StatusConflict, structs.Message{Message: "User already exists!"})
	}
}

func getUser(c *gin.Context) {
	/*
		Gets an existing user from the 'users' collection in the database

		@param: user id via api path parameter
		@return: user structure
	*/
	userId := c.Param("id")

	// Gets the 'users' collection from the database
	database := client.Database("VaasDatabase")
	userCollection := database.Collection("users")

	// Aggregate matching pipeline query
	var existingUser structs.User
	criteria := map[string]interface{}{
		"id": userId,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := userCollection.FindOne(ctx, criteria).Decode(&existingUser)

	// If no documents were found, insert the new user into the collection and return a corresponding confirmation message. Otherwise, return a stale message
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, structs.Message{Message: "No documents found"})
	} else {
		c.JSON(http.StatusOK, existingUser)
	}
}

func updateUser(c *gin.Context) {
	/*
		Updates an existing user from the 'users' collection in the database

		@param: username via api path parameter
		@return: JSON confirmation message
	*/
	userId := c.Param("userId")

	// Gets the HTTP header and body
	userUpdateData := &structs.UserUpdate{}
	if err := c.ShouldBindJSON(userUpdateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update document in database
	database := client.Database("VaasDatabase")
	gameCollection := database.Collection("users")
	filter := bson.D{
		{Key: "id", Value: userId},
	}
	update := bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "numgames", Value: userUpdateData.ChangeInNumGames},
			{Key: "totalguesses", Value: userUpdateData.ChangeInTotalGuesses},
		}},
	}

	_, err := gameCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update game in mongo: " + err.Error()})
		return
	}

	// Return confirmation message
	c.JSON(http.StatusOK, &structs.Message{
		Message: "user updated successfully",
	})
}
