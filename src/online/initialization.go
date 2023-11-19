package main

import (
	"net/http"
	"vaas/structs"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func newMultiplayerGame(c *gin.Context) {
	/*
		This API method takes in a metadata struct in the request body, sends it to Engine to make a new game, creates a MultiplayerGame struct, and sends it to Mongo to store

		@param: metadata for new game (structs.Metadata)
		@return: new multiplayer game created (structs.MultiplayerGame)
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
		Create multiplayer game structure. The game ID and word received in Engine response are put into the 'Games' map and 'Word' field, respectively
	*/
	multiplayerGame := structs.MultiplayerGame{
		MultiplayerGameID: uuid.NewString(),
		HostID:            newMetadata.UserId,
		Games:             make(map[string]string),
		State:             "waiting",
		WinnerID:          "",
		Word:              newGame.Word,
	}
	multiplayerGame.Games[newMetadata.UserId] = newGame.Metadata.GameID

	// Send multiplayer game to Mongo communication method
	if err = mongo_createMultiplayerGame(multiplayerGame); err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
	} else {
		c.JSON(http.StatusOK, multiplayerGame)
	}
}

func getMultiplayerGame(c *gin.Context) {
	/*
		Takes in a path parameter for the multiplayer game ID, and communicates with Mongo to return the multiplayer game struct associated with it

		@param: multiplayer game id in path parameter (string)
		@return: multiplayer game (structs.MultiplayerGame)
	*/
	id := c.Param("id")
	game, err := helper_getMultiplayerGame(id)
	if err != nil {
		c.JSON(http.StatusNotFound, structs.Message{Message: "Online: Game could not be found"})
		return
	}
	c.JSON(http.StatusOK, game)
}

func helper_getMultiplayerGame(multiplayerGameID string) (*structs.MultiplayerGame, error) {
	/*
		Takes in a path parameter for the multiplayer game ID, and communicates with Mongo to return the multiplayer game struct associated with it

		@param: multiplayer game id in path parameter (string)
		@return: multiplayer game (structs.MultiplayerGame)
	*/

	multiplayerGame, err := mongo_getMultiplayerGame(multiplayerGameID)
	if err != nil {
		return nil, err
	}
	return multiplayerGame, nil
}

func joinMultiplayerGame(c *gin.Context) {
	/*
		Takes in a metadata struct as well as a multiplayer game ID path parameter. Gets the multiplayer game from Mongo, sets the enforced word of a new game to the word of the multiplayer game, and adds the game/user to the multiplayer game

		@param: metadata for new game (structs.Metadata) as well as the multiplayer game ID (string) as a path parameter
		@return: updated multiplayer game (structs.MultiplayerGame)
	*/
	id := c.Param("id")
	multiplayerGame, err := helper_getMultiplayerGame(id)
	if err != nil {
		c.JSON(http.StatusNotFound, structs.Message{Message: "Online: Game could not be found"})
		return
	}

	newMetadata := structs.GameMetadata{}
	if err := c.ShouldBindJSON(&newMetadata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request format, expected {wordLength: int, maxGuesses: int}"})
		return
	}

	// Add enforced word to game metadata
	newMetadata.EnforcedWord = multiplayerGame.Word

	// Send metadata to engine to create new game
	newGame, err := engine_newGame(newMetadata)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}

	// Add new user to the field
	multiplayerGame.Games[newMetadata.UserId] = newGame.Metadata.GameID

	// Send multiplayer game ID, new individual game ID, and user ID to add to new multiplayer game in Mongo
	if err = mongo_addUserToMultiplayerGame(multiplayerGame.MultiplayerGameID, newGame); err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}

	// Respond based on Mongo response
	c.JSON(http.StatusOK, multiplayerGame)
}
