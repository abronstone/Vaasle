package main

import (
	"bytes"
	"encoding/json"
	"io"
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

	bodyBytes, err := json.Marshal(newGameCustomMetadata)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal the request body: " + err.Error()})
		return
	}
	bodyBuffer := bytes.NewBuffer(bodyBytes)

	res, err := http.Post("http://online:8000/newMultiplayerGame", "application/json", bodyBuffer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}
	defer res.Body.Close()

	game := structs.MultiplayerGame{}

	responseBodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}

	err = json.Unmarshal(responseBodyBytes, &game)
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

	// Create a new HTTP client
	client := &http.Client{}
	endpoint := "http://online:8000/joinMultiplayerGame/" + id

	// Convert struct to JSON
	jsonData, err := json.Marshal(newGameCustomMetadata)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}

	// Create a new request
	req, err := http.NewRequest(http.MethodPut, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}

	// Set Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Execute the request
	res, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: "could not join game"})
		return
	}

	responseBodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}

	game := &structs.MultiplayerGame{}
	err = json.Unmarshal(responseBodyBytes, game)
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
	client := &http.Client{}
	endpoint := "http://online:8000/startMultiplayerGame/" + id

	// Create a new request
	req, err := http.NewRequest(http.MethodPut, endpoint, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}

	// Set Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Execute the request
	res, err := client.Do(req)
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
	client := &http.Client{}
	endpoint := "http://online:8000/refreshMultiplayerGame/" + id

	// Create a new request
	req, err := http.NewRequest(http.MethodPut, endpoint, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}

	// Set Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Execute the request
	res, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: "could not refresh game"})
		return
	}

	responseBodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}

	update := &structs.MultiplayerFrontendUpdate{}
	err = json.Unmarshal(responseBodyBytes, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, update)
}
