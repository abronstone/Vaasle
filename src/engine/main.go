package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Map containing all Wordle game instances.
var games map[int]*game = make(map[int]*game, 0)

// The metadata of a Wordle game.
type gameMetadata struct {
	GameID     int `json:"gameID"`
	WordLength int `json:"wordLength"`
	MaxGuesses int `json:"maxGuesses"`
}

// A Wordle game.
type game struct {
	Metadata gameMetadata `json:"metadata"`
	Guesses  []string     `json:"guesses"`
}

// Returns a settings configuration for a new game: gameID, wordLength, maxGuesses.
func getNewGameConfiguration(gameIDGetter func() int, c *gin.Context) (int, int, int) {
	gameID := gameIDGetter()
	wordLength := 5
	maxGuesses := 6

	if num, err := strconv.Atoi(c.Query("wordLength")); err == nil {
		wordLength = num
	}
	if num, err := strconv.Atoi(c.Query("maxGuesses")); err == nil {
		maxGuesses = num
	}

	return gameID, wordLength, maxGuesses
}

// Creates a new instance of a Wordle game.
// Returns JSON information about the newly instantiated game.
func newGame(c *gin.Context) {
	// Eventually, we want to replace this naive function.
	// To get a unique game ID, we should be querying MongoDB for a unique ID.
	naiveGetUniqueGameID := func() int {
		return int(uuid.New().ID())
	}

	gameID, wordLength, maxGuesses := getNewGameConfiguration(naiveGetUniqueGameID, c)

	newGame := &game{
		Metadata: gameMetadata{
			GameID:     gameID,
			WordLength: wordLength,
			MaxGuesses: maxGuesses,
		},
		Guesses: make([]string, 0),
	}

	games[newGame.Metadata.GameID] = newGame
	c.JSON(http.StatusOK, newGame)
}

// Default home endpoint for checking if the engine is running.
func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Engine is running"})
}

// Engine API's main method.
func main() {
	router := gin.Default()

	router.GET("/", home)
	router.GET("/newgame", newGame)

	router.Run("0.0.0.0:5001")
}
