package main

import (
	"context"
	"net/http"
	"vaas/structs"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func newGame(c *gin.Context) {
	/*
		Creates a new game structure with the specified parameters, gets a random word that conforms to the word length (soon to be language), and returns the word

		OVERWRITE SAME GAME IF IT EXISTS

		@param: game metadata structure in HTTP body
		@return: word
	*/

	// Get metadata from HTTP body
	var metadata structs.GameMetadata
	if err := c.ShouldBindJSON(&metadata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gameID := metadata.GameID
	wordLength := metadata.WordLength
	maxGuesses := metadata.MaxGuesses
	dateCreated := metadata.DateCreated
	userName := metadata.UserName
	gameMetadata := structs.GameMetadata{GameID: gameID, WordLength: wordLength, MaxGuesses: maxGuesses, DateCreated: dateCreated, UserName: userName}

	// Get collections
	database := client.Database("VaasDatabase")
	wordCollection := database.Collection("words")
	gameCollection := database.Collection("games")

	deleteFilter := bson.D{{"metadata.gameid", bson.D{{"$eq", gameID}}}}
	gameCollection.DeleteOne(context.TODO(), deleteFilter)

	// Get a random word
	var randomWord bson.M
	cursor, err := wordCollection.Aggregate(context.TODO(), bson.A{
		bson.D{{"$match", bson.D{{"solution", true}, {"length", wordLength}}}},
		bson.D{{"$sample", bson.D{{"size", 1}}}},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve a word from mongo: " + err.Error()})
		return
	}

	if cursor.Next(context.TODO()) {
		if err := cursor.Decode(&randomWord); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decode word from mongo: " + err.Error()})
			return
		}
	}
	defer cursor.Close(context.Background())

	// Create new game structure and insert into database
	guesses := [][2]string{}
	game := structs.Game{Word: randomWord["word"].(string), Metadata: gameMetadata, Guesses: guesses, State: "ongoing"}
	gameCollection.InsertOne(context.TODO(), game)

	// Return initialized game state
	c.JSON(http.StatusOK, game)
}

func updateGameState(c *gin.Context) {
	/*
		Updates the game state and guesses

		@param: updated game structure in HTTP body
		@return: JSON confirmation message
	*/

	// Gets the HTTP header and body
	var gameData structs.Game
	if err := c.ShouldBindJSON(&gameData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newState := gameData.State
	newGuesses := gameData.Guesses
	gameID := gameData.Metadata.GameID

	// Update document in database
	database := client.Database("VaasDatabase")
	gameCollection := database.Collection("games")
	filter := bson.D{{"metadata.gameid", gameID}}
	update := bson.D{{"$set", bson.D{{"state", newState}, {"guesses", newGuesses}}}}

	_, err := gameCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update game in mongo: " + err.Error()})
		return
	}

	// Return confirmation message
	c.JSON(http.StatusOK, &structs.Message{
		Message: "game updated successfully",
	})
}

func getGame(c *gin.Context) {
	/*
		Finds and returns the specified game from the database

		@param: game ID
		@return: game structure with the associated game ID
	*/

	// Get game ID from path parameters
	gameID := c.Param("id")

	database := client.Database("VaasDatabase")
	gameCollection := database.Collection("games")

	// Create match pipeline stage
	matchStage := bson.A{
		bson.D{{"$match", bson.D{{"metadata.gameid", gameID}}}},
	}
	var game structs.Game
	// Run aggregation
	cursor, err := gameCollection.Aggregate(context.TODO(), matchStage)
	defer cursor.Close(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve game from mongo: " + err.Error()})
		return
	}
	if cursor.Next(context.TODO()) {
		if err := cursor.Decode(&game); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decode game from mongo: " + err.Error()})
			return
		}
	}

	// Return game
	c.JSON(http.StatusOK, game)
}
