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
Checks the availability of the engine container.
Performs an HTTP GET request to the engine's health check endpoint.
If the engine is unreachable or returns an error, respond with a 500 Internal Server Error.
Otherwise, return the engine's response.

@param: None (uses Gin context)
@return: Response from the engine or an error message
*/
func api_pingEngine(c *gin.Context) {
	res, err := http.Get("http://engine:5001/")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ping to engine failed: " + err.Error()})
		return
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failure decoding response from engine: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, string(bodyBytes))
}

/*
Creates a new game instance.
Sends a POST request to the engine's 'newGame' endpoint with custom game metadata.
On success, returns the initial state of the new game.
Handles errors by returning appropriate HTTP status codes and error messages.

@param: Custom game metadata in the request body
@return: Initial game state or an error message
*/
func api_newGame(c *gin.Context) {
	newGameCustomMetadata := structs.GameMetadata{}

	// Bind the incoming JSON body to the newGameCustomMetadata struct
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

	// Call the engine's newGame endpoint
	res, err := http.Post("http://engine:5001/newGame", "application/json", bodyBuffer)

	// If the engine is down, return an error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a new game: " + err.Error()})
		return
	}

	defer res.Body.Close()

	// Create a newGame variable and unmarshal the response body into it
	newGame := structs.Game{}
	responseBodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed read response body from engine: " + err.Error()})
		return
	}

	err = json.Unmarshal(responseBodyBytes, &newGame)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal response body from engine: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, newGame.GetShareable())
}

/*
Retrieves the current state of an existing game.
Extracts the game ID from the URL, makes a GET request to the engine's 'getGame' endpoint.
Returns the game's state.
Handles error scenarios such as an unreachable engine or unprocessable game ID by returning appropriate HTTP status codes and error messages.

@param: Game ID from the URL path parameter
@return: Current game state or an error message
*/
func api_getGame(c *gin.Context) {
	// Get the gameID from the URL
	gameID := c.Param("id")

	// Call the engine's getGame endpoint
	res, err := http.Get("http://engine:5001/getGame/" + gameID)

	// If the engine is down, return an error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get game from engine: " + err.Error()})
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get game from engine: " + err.Error()})
		return
	}

	// Create a currentGame variable and unmarshal the response body into it
	currentGame := structs.Game{}
	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed read response body from engine: " + err.Error()})
		return
	}

	err = json.Unmarshal(bodyBytes, &currentGame)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal response body from engine" + err.Error()})
		return
	}

	if currentGame.Metadata.GameID == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get game from engine: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, currentGame.GetShareable())
}

/*
Processes a player's guess for a game.
Receives a guess in the request body, forwards it to the engine's 'makeGuess' endpoint.
Returns the updated game state.
Handles error cases like malformed requests or engine errors, responding with corresponding HTTP status codes and error messages.

@param: Player's guess in the request body
@return: Updated game state or an error message
*/
func api_makeGuess(c *gin.Context) {
	// Define the format of the request body
	guess := structs.Guess{}

	// Bind the incoming JSON body to the guess struct
	if err := c.ShouldBindJSON(&guess); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request format, expected {id: string, guess: string}"})
		return
	}

	// Use json.Marshal to convert the guess object to a JSON-formatted []byte
	bodyBytes, err := json.Marshal(guess)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal the request body: " + err.Error()})
		return
	}

	// Convert []byte to io.Reader
	bodyBuffer := bytes.NewBuffer(bodyBytes)

	// Call the engine's makeGuess endpoint
	res, err := http.Post("http://engine:5001/makeGuess", "application/json", bodyBuffer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make guess, please enter a valid 5 letter english word: " + err.Error()})
		return
	}

	if res.StatusCode != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make guess, please enter a valid 5 letter english word"})
		return
	}

	defer res.Body.Close()

	// Create a currentGame variable and unmarshal the response body into it
	currentGame := structs.Game{}
	bodyBytes, err = io.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed read response body from engine: " + err.Error()})
		return
	}

	if err := json.Unmarshal(bodyBytes, &currentGame); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal response body from engine: " + err.Error()})
		return
	}

	if currentGame.Metadata.GameID == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get game from engine: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, currentGame.GetShareable())
}
