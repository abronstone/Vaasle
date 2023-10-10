package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type game struct {
	Metadata gameMetadata `json:"metadata"`
	Guesses  [][2]string  `json:"guesses"`
	State    string       `json:"state"`
	Word     string       `json:"-"`
}

// The metadata of a Wordle game.
type gameMetadata struct {
	GameID     string `json:"gameID"`
	WordLength int    `json:"wordLength"`
	MaxGuesses int    `json:"maxGuesses"`
}

func main() {
	router := gin.Default()

	// Add CORS middleware
	router.Use(cors.Default())

	router.GET("/", api_home)
	router.GET("/pingEngine", api_pingEngine)
	router.POST("/newGame", api_newGame)
	// router.GET("/getGame/:id", api_getGame)
	// router.POST("/makeGuess", api_makeGuess)

	router.Run("0.0.0.0:5001")
}

func api_home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, map[string]string{"message": "Play game is running"})
}

func api_pingEngine(c *gin.Context) {
	res, err := http.Get("http://engine:5001/")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Ping to engine failed"})
		return
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Ping to engine failed"})
		return
	}

	c.JSON(http.StatusOK, string(bodyBytes))
}

func home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, map[string]string{"message": "Play game is running"})
}

// Calls the appropriate endpoint in the engine to make a new game and returns the game's public state as JSON.
func api_newGame(c *gin.Context) {
	// Call the engine's newGame endpoint
	res, err := http.Post("http://engine:5000/newGame", "application/json", c.Request.Body)

	// If the engine is down, return an error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a new game"})
		return
	}

	defer res.Body.Close()

	// Create a newGame variable and unmarshal the response body into it
	newGame := game{}
	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed read response body"})
		return
	}

	err = json.Unmarshal(bodyBytes, &newGame)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal response body"})
		return
	}

	c.JSON(http.StatusOK, newGame)

}
