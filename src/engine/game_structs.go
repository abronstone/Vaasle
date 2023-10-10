package main

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// The metadata of a Wordle game.
type gameMetadata struct {
	GameID     string `json:"gameID"`
	WordLength int    `json:"wordLength"`
	MaxGuesses int    `json:"maxGuesses"`
}

// Create a new gameMetadata struct.
func newGameMetadata(c *gin.Context) *gameMetadata {
	defaultWordLength := 5
	defaultMaxGuesses := 6

	metadata := gameMetadata{
		WordLength: defaultWordLength,
		MaxGuesses: defaultMaxGuesses,
	}

	// Use default values if we fail to parse the POST request body.
	if err := c.ShouldBindJSON(&metadata); err != nil {
		metadata.WordLength = defaultWordLength
		metadata.MaxGuesses = defaultMaxGuesses
	}

	metadata.GameID = uuid.NewString()
	return &metadata
}

// A Wordle game.
type game struct {
	Metadata gameMetadata `json:"metadata"`
	Guesses  [][2]string  `json:"guesses"`
	State    string       `json:"state"`
	Word     string       `json:"-"`
}

// Create a new game struct.
func newGame(c *gin.Context) *game {
	metadata := newGameMetadata(c)
	return &game{
		Metadata: *metadata,
		Guesses:  make([][2]string, 0, metadata.MaxGuesses),
		State:    "ongoing",
	}
}

// Set the secret word of a game (only allowed at game initialization).
func (g *game) setWord(word string) error {
	if g.Word != "" {
		return errors.New("game already has a word")
	}
	if len(word) != g.Metadata.WordLength {
		return fmt.Errorf(`given word "%s" is not of length %d`, word, g.Metadata.WordLength)
	}
	g.Word = word
	return nil
}
