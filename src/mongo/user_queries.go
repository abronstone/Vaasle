package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"vaas/structs"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func newUser(c *gin.Context) {
	/*
		Adds a new user to the 'users' collection in the database

		@param: username via api path parameter
		@return: confirmation message
	*/
	// Creates a new user with the specified username

	var user structs.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database := client.Database("VaasDatabase")
	userCollection := database.Collection("users")

	// Execute the query
	cursor, err := userCollection.Find(context.TODO(), bson.M{"username": user.UserName})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	defer cursor.Close(context.TODO())
	// If no documents were found, insert the new user into the collection and return a corresponding confirmation message. Otherwise, return a stale message
	if !cursor.Next(context.TODO()) {
		userCollection.InsertOne(context.TODO(), user)
		c.JSON(http.StatusOK, structs.Message{Message: "User inserted into database"})
	} else {
		c.JSON(http.StatusNotFound, structs.Message{Message: "User already exists!"})
	}
}

func getUser(c *gin.Context) {
	/*
		Gets an existing user from the 'users' collection in the database

		@param: username via api path parameter
		@return: user structure
	*/
	userName := c.Param("username")

	// Gets the 'users' collection from the database
	database := client.Database("VaasDatabase")
	userCollection := database.Collection("users")

	// Aggregate matching pipeline query
	var existingUser structs.User
	criteria := map[string]interface{}{
		"username": userName,
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
