package main

import (
	"net/http"
	"vaas/structs"

	"github.com/gin-gonic/gin"
)

/*
Creates a new multiplayer game instance.
Sends a POST request to the 'newMultiplayerGame' endpoint with custom game metadata.
On success, returns the initial state of the new multiplayer game.
Handles errors by returning appropriate HTTP status codes and error messages.

@param: Custom game metadata in the request body
@return: Initial multiplayer game state or an error message
*/
func api_createMultiplayerGame(c *gin.Context) {
	newGameCustomMetadata := structs.GameMetadata{}
	if err := c.ShouldBindJSON(&newGameCustomMetadata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request format, expected {wordLength: int, maxGuesses: int, userID: string}"})
		return
	}

	game := structs.MultiplayerGame{}
	// Make POST request, encoding 'newGameCustomMetadata' to request body, decode response body into 'game' of type 'structs.Game'
	_, err := structs.MakePostRequest[structs.MultiplayerGame]("http://online:8000/newMultiplayerGame", newGameCustomMetadata, &game)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, game.GetShareable())
}

/*
Joins an existing multiplayer game.
Retrieves the game ID from the URL, sends a PUT request to the 'joinMultiplayerGame' endpoint with the provided metadata.
On successful joining, returns the current state of the multiplayer game.
Handles errors by returning appropriate HTTP status codes and error messages.

@param: Game ID from the URL and custom game metadata in the request body
@return: Current multiplayer game state or an error message
*/
func api_joinMultiplayerGame(c *gin.Context) {
	id := c.Param("id")

	newGameCustomMetadata := structs.GameMetadata{}
	if err := c.ShouldBindJSON(&newGameCustomMetadata); err != nil {
		c.JSON(http.StatusBadRequest, structs.Message{Message: "Bad request format, expected {wordLength: int, maxGuesses: int, userID: string}"})
		return
	}

	game := structs.Game{}
	// Make PUT request, encoding 'newGameCustomMetadata' to request body, decode response body into 'game' of type 'structs.Game'
	_, err := structs.MakePutRequest[structs.Game]("http://online:8000/joinMultiplayerGame/"+id, newGameCustomMetadata, &game)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, game.GetShareable())
}

/*
Starts an already created multiplayer game.
Retrieves the game ID from the URL, sends a PUT request to the 'startMultiplayerGame' endpoint.
Confirms the start of the game by returning a success message.
Handles errors by returning appropriate HTTP status codes and error messages.

@param: Game ID from the URL
@return: Confirmation message or an error message
*/
func api_startMultiplayerGame(c *gin.Context) {
	id := c.Param("id")

	// Make PUT request, no encoding/decoding needed
	res, err := structs.MakePutRequest[any]("http://online:8000/startMultiplayerGame/"+id, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: "could not start game"})
		return
	}

	c.JSON(http.StatusOK, structs.Message{Message: "started game"})
}

/*
Refreshes the state of an existing multiplayer game.
Retrieves the game ID from the URL, sends a PUT request to the 'refreshMultiplayerGame' endpoint.
Returns the latest state of the multiplayer game.
Handles errors by returning appropriate HTTP status codes and error messages.

@param: Game ID from the URL
@return: Updated multiplayer game state or an error message
*/
func api_refreshMultiplayerGame(c *gin.Context) {
	id := c.Param("id")
	update := structs.MultiplayerFrontendUpdate{}
	// Make PUT request, no request body, decode response body into 'update' of type 'structs.MultiplayerFrontendUpdate'
	res, err := structs.MakePutRequest[structs.MultiplayerFrontendUpdate]("http://online:8000/refreshMultiplayerGame/"+id, nil, &update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: "could not refresh game"})
		return
	}

	c.JSON(http.StatusOK, update)
}
