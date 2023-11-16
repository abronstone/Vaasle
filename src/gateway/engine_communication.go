package main

import (
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

	newGame := structs.Game{}

	res, err := structs.MakePostRequest[structs.Game]("http://engine:5001/newGame", newGameCustomMetadata, &newGame)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "newGame post request failed " + err.Error()})
		return
	}
	defer res.Body.Close()

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
	currentGame := structs.Game{}
	// Make GET request, decode response into 'currentGame' of type 'structs.Game'
	_, err := structs.MakeGetRequest[structs.Game]("http://engine:5001/getGame/"+gameID, &currentGame)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get game from engine: " + err.Error()})
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

	currentGame := structs.Game{}
	// Make POST request, encoding 'guess' to request body, decode response body into 'currentGame' of type 'structs.Game'
	res, err := structs.MakePostRequest[structs.Game]("http://engine:5001/makeGuess", guess, &currentGame)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error making guess, make sure you entered a valid input" + err.Error()})
		return
	}
	defer res.Body.Close()

	if currentGame.Metadata.GameID == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get game from engine: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, currentGame.GetShareable())
}
