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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request format, expected {wordLength: int, maxGuesses: int}"})
		return
	}

	// Send metadata to engine to create new game
	newGame, err := engine_newGame(newMetadata)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: "Engine newGame issue: " + err.Error()})
	}

	/*
		Create shared game structure. The game ID and word received in Engine response are put into the 'Games' map and 'Word' field, respectively
	*/
	sharedGame := structs.SharedGame{
	    SharedGameID: uuid.NewString(),
	    HostID: newMetadata.UserId,
	    Games: make(map[string]string),
	    State: "waiting",
	    WinnerID: "",
	    Word: newGame.Word,
	}
	sharedGame.Games[newMetadata.UserId] = newGame.Metadata.GameID

	// Send shared game to Mongo communication method
	mongo_response, err := mongo_createSharedGame(sharedGame)

	// Respond based on Mongo response
	if err != nil || mongo_response == "Could not create shared game due to Mongo error" {
		c.JSON(http.StatusBadRequest, structs.Message{Message: err.Error()})
	} else {
		c.JSON(http.StatusOK, sharedGame)
	}
}

func getSharedGame(c *gin.Context) {
	c.JSON(http.StatusOK, structs.Message{Message: "getSharedGame working"})
}

// Needs to enforce the word when calling "/newGame/" from engine
func joinSharedGame(c *gin.Context) {
	/*
		1. Get shared game from Mongo
		2. Create enforced word game
		3. Send update to Mongo to add game to shared game
	*/
	c.JSON(http.StatusOK, structs.Message{Message: "joinSharedGame working"})
}
