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
	router.PUT("/createUser", api_newUser)
	router.PUT("/login/:username", api_login)

	router.Run("0.0.0.0:5001")
}

func api_home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, map[string]string{"message": "Gateway is running"})
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

func api_newUser(c *gin.Context) {
	// Define a struct to bind the request body

	type NewUserRequestBody struct {
		UserName string `json:"userName"`
		Id       string `json:"id"`
	}

	var requestBody NewUserRequestBody

	// Bind JSON to struct
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Convert struct to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error converting data to JSON"})
		return
	}

	// Create a new request
	req, err := http.NewRequest(http.MethodPut, "http://online:80/create-user", bytes.NewBuffer(jsonData))

	// Handle request creation error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a new request: " + err.Error()})
		return
	}

	// Set Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Execute the request
	res, err := client.Do(req)

	// Handle execution error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute user creation request: " + err.Error()})
		return
	}

	// Close the response body
	defer res.Body.Close()

	// Check for status codes
	if res.StatusCode == http.StatusBadRequest {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	if res.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func api_login(c *gin.Context) {
	var userName string = c.Param("username")

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new request
	req, err := http.NewRequest(http.MethodPut, "http://online:80/login/"+userName, nil)

	// Handle request creation error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a new request: " + err.Error()})
		return
	}

	// Execute the request
	res, err := client.Do(req)

	// Handle execution error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute login request: " + err.Error()})
		return
	}

	// Close the response body
	defer res.Body.Close()

	// Check for status codes
	if res.StatusCode == http.StatusUnauthorized {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
		return
	}

	if res.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error logging in"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
