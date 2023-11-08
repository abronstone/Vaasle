package main

import (
	"net/http"
	"vaas/structs"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func newSharedGame(c *gin.Context) {
	/*
		This API method takes in a metadata struct in the request body, sends it to Engine to make a new game, creates a SharedGame struct, and sends it to Mongo to store

		@param: metadata for new game (structs.Metadata)
		@return: new shared game created (structs.SharedGame)
	*/
	newMetadata := structs.GameMetadata{}
	if err := c.ShouldBindJSON(&newMetadata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request format, expected {wordLength: int, maxGuesses: int, userID: string}"})
		return
	}

	// Send metadata to engine to create new game
	newGame, err := engine_newGame(newMetadata)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: "Engine newGame issue: " + err.Error()})
		return
	}

	/*
		Create shared game structure. The game ID and word received in Engine response are put into the 'Games' map and 'Word' field, respectively
	*/
	sharedGame := structs.SharedGame{
		SharedGameID: uuid.NewString(),
		HostID:       newMetadata.UserId,
		Games:        make(map[string]string),
		State:        "waiting",
		WinnerID:     "",
		Word:         newGame.Word,
	}
	sharedGame.Games[newMetadata.UserId] = newGame.Metadata.GameID

	// Send shared game to Mongo communication method
	if err = mongo_createSharedGame(sharedGame); err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
	} else {
		c.JSON(http.StatusOK, sharedGame)
	}
}

func getSharedGame(c *gin.Context) *structs.SharedGame {
	/*
		Takes in a path parameter for the shared game ID, and communicates with Mongo to return the shared game struct associated with it

		@param: shared game id in path parameter (string)
		@return: shared game (structs.SharedGame)
	*/
	sharedGameID := c.Param("id")

	sharedGame, err := mongo_getSharedGame(sharedGameID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: "ERROR: " + err.Error()})
		return nil
	}

	c.JSON(http.StatusOK, sharedGame)
	return sharedGame
}

func joinSharedGame(c *gin.Context) {
	/*
		Takes in a metadata struct as well as a shared game ID path parameter. Gets the shared game from Mongo, sets the enforced word of a new game to the word of the shared game, and adds the game/user to the shared game

		@param: metadata for new game (structs.Metadata) as well as the shared game ID (string) as a path parameter
		@return: updated shared game (structs.SharedGame)
	*/
	sharedGame := getSharedGame(c)
	if sharedGame == nil {
		c.JSON(http.StatusNotFound, structs.Message{Message: "Online: Game could not be found"})
	}

	newMetadata := structs.GameMetadata{}
	if err := c.ShouldBindJSON(&newMetadata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request format, expected {wordLength: int, maxGuesses: int}"})
		return
	}

	// Add enforced word to game metadata
	newMetadata.EnforcedWord = sharedGame.Word

	// Send metadata to engine to create new game
	newGame, err := engine_newGame(newMetadata)

	// Add new user to the field
	sharedGame.Games[newMetadata.UserId] = newGame.Metadata.GameID

	// Send shared game ID, new individual game ID, and user ID to add to new shared game in Mongo
	if err = mongo_addUserToSharedGame(sharedGame.SharedGameID, newGame.GameID, newMetadata.UserId); err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}

	// Respond based on Mongo response
	c.JSON(http.StatusOK, sharedGame)
}
