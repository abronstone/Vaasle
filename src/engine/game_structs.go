package main

import (
	"errors"
	"fmt"
	"strconv"

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
	wordLength := 5
	maxGuesses := 6

	if num, err := strconv.Atoi(c.Query("wordLength")); err == nil {
		wordLength = num
	}
	if num, err := strconv.Atoi(c.Query("maxGuesses")); err == nil {
		maxGuesses = num
	}

	return &gameMetadata{
		GameID:     uuid.NewString(),
		WordLength: wordLength,
		MaxGuesses: maxGuesses,
	}
}

// A Wordle game.
type game struct {
	Metadata gameMetadata `json:"metadata"`
	Guesses  [][2]string  `json:"guesses"`
	State    string       `json:"state"`
	word     string
}

// Create a new game struct.
func newGame(c *gin.Context) *game {
	return &game{
		Metadata: *newGameMetadata(c),
		Guesses:  make([][2]string, 0),
		State:    "ongoing",
	}
}

// Set the secret word of a game (only allowed at game initialization).
func (g *game) setWord(word string) error {
	if g.word != "" {
		return errors.New("game already has a word")
	}
	if len(word) != g.Metadata.WordLength {
		return fmt.Errorf(`given word "%s" is not of length %d`, word, g.Metadata.WordLength)
	}
	g.word = word
	return nil
}
