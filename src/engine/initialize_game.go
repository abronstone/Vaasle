package main

import (
	"errors"
	"fmt"
	"vaas/structs"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Create a new gameMetadata struct.
func newGameMetadata(c *gin.Context) *structs.GameMetadata {
	defaultWordLength := 5
	defaultMaxGuesses := 6

	metadata := structs.GameMetadata{}

	// Use default values if we fail to parse the POST request body.
	if err := c.ShouldBindJSON(&metadata); err != nil {
		metadata.WordLength = defaultWordLength
		metadata.MaxGuesses = defaultMaxGuesses
	}

	metadata.GameID = uuid.NewString()
	return &metadata
}

// Create a new game struct.
func newGame(c *gin.Context) *structs.Game {
	metadata := newGameMetadata(c)
	return &structs.Game{
		Metadata: *metadata,
		Guesses:  make([][2]string, 0, metadata.MaxGuesses),
		State:    "ongoing",
	}
}

// Set the secret word of a game (only allowed at game initialization).
func setWord(g *structs.Game, word string) error {
	if g.Word != "" {
		return errors.New("game already has a word")
	}
	if len(word) != g.Metadata.WordLength {
		return fmt.Errorf(`given word "%s" is not of length %d`, word, g.Metadata.WordLength)
	}
	g.Word = word
	return nil
}
