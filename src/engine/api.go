package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"vaas/structs"

	"github.com/gin-gonic/gin"
)

// Engine API's main method.
func main() {
	router := gin.Default()

	router.GET("/", api_home)
	router.POST("/newGame", api_newGame)
	router.GET("/getGame/:id", api_getGame)
	router.POST("/makeGuess", api_makeGuess)
	router.GET("/pingPlayGame", api_pingPlayGame)

	router.GET("/getGameExposed/:id", api_getGameExposed)

	router.Run("0.0.0.0:5001")
}

// Default home endpoint for checking if the engine is running.
func api_home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "engine is running"})
}

// Creates a new Wordle game and returns its public state as a JSON object.
func api_newGame(c *gin.Context) {
	newGame := newGame(c)

	word, err := mongo_submitNewGame(newGame.Metadata)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve a word from mongo: " + err.Error()})
		return
	}

	err = setWord(newGame, word)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to initialize game"})
		return
	}

	registerGame(newGame)
	c.JSON(http.StatusOK, newGame)
}

// Returns the game struct with the specified ID as a JSON object.
func api_getGame(c *gin.Context) {
	if game, err := getGame(c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "could not find game"})
	} else {
		c.JSON(http.StatusOK, game)
	}
}

// Returns the game struct with the specified ID as a JSON object.
// Exposes the word, only used for debugging.
func api_getGameExposed(c *gin.Context) {
	if game, err := getGame(c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "could not find game"})
	} else {
		c.JSON(http.StatusOK, game.ConvertToGameExposed())
	}
}

// Returns the game struct with the specified ID as a JSON object.
func api_makeGuess(c *gin.Context) {
	requestBody := structs.Guess{}
	c.ShouldBindJSON(&requestBody)

	game, err := getGame(requestBody.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "could not find game"})
		return
	}

	err = makeGuess(game, requestBody.Guess)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = mongo_updateGame(game)
	if err != nil {
		log.Println(err.Error()) // not a fatal error, we can just log it
	}

	c.JSON(http.StatusOK, game)
}

// Pings the play-game endpoint and forwards its response.
func api_pingPlayGame(c *gin.Context) {
	res, err := http.Get("http://play-game:5001/")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ping to play-game failed"})
		return
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failure to parse play-game response body: " + err.Error()})
		return
	}

	result := struct {
		Message string `json:"message"`
	}{}
	json.Unmarshal(bodyBytes, &result)
	c.JSON(http.StatusOK, &result)
}
