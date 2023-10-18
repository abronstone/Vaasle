package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"vaas/structs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Add CORS middleware
	router.Use(cors.Default())

	router.GET("/", api_home)
	router.GET("/pingEngine", api_pingEngine)
	router.POST("/newGame", api_newGame)
	router.GET("/getGame/:id", api_getGame)
	router.POST("/makeGuess", api_makeGuess)

	router.Run("0.0.0.0:5001")
}

func api_home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, map[string]string{"message": "Play game is running"})
}

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

// Calls the appropriate endpoint in the engine to make a new game
// and returns the game's public state as JSON.
// Takes in an CustomMetaData and in the request body 
// and returns a game struct.
func api_newGame(c *gin.Context) {
	newGameCustomMetadata := structs.GameMetadata{}

	// Bind the incoming JSON body to the newGameCustomMetadata struct
	if err := c.ShouldBindJSON(&newGameCustomMetadata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request format, expected {wordLength: int, maxGuesses: int}"})
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

	c.JSON(http.StatusOK, newGame)
}

// Calls the appropriate endpoint in the engine to retrieve an existing game
// and returns the game's public state as JSON.
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

	c.JSON(http.StatusOK, currentGame)
}

// Calls the appropriate endpoint in the engine to make a guess
// and returns the game's public state as JSON.
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make guess with engine: " + err.Error()})
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

	c.JSON(http.StatusOK, currentGame)
}
