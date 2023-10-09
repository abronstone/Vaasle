package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Engine API's main method.
func main() {
	router := gin.Default()

	router.GET("/", api_home)
	router.GET("/newGame", api_newGame)
	router.GET("/pingPlayGame", api_pingPlayGame)

	router.Run("0.0.0.0:5001")
}

// Default home endpoint for checking if the engine is running.
func api_home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Engine is running"})
}

// Creates a new Wordle game and returns its public state as a JSON object.
func api_newGame(c *gin.Context) {
	newGame := newGame(c)

	word, err := submitNewGame(newGame.Metadata.GameID, newGame.Metadata.WordLength)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "failed to retrieve a word"})
		return
	}

	err = newGame.setWord(word)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "failed to initialize game"})
		return
	}

	registerGame(newGame)
	c.JSON(http.StatusOK, newGame)
}

// Pings the play-game endpoint and forwards its response.
func api_pingPlayGame(c *gin.Context) {
	res, err := http.Get("http://play-game:5001/")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Ping to play-game failed"})
		return
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Ping to play-game failed"})
		return
	}

	result := struct {
		Message string `json:"message"`
	}{}
	json.Unmarshal(bodyBytes, &result)
	c.JSON(http.StatusOK, &result)
}
