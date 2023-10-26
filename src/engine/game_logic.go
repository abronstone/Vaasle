package main

import (
	"fmt"
	"strings"
	"vaas/structs"
)

// Gets the game with the specified ID.
// Queries the Mongo API if not already present in the engine's cache.
func getGame(id string) (*structs.Game, error) {
	if game, ok := games.load(id); ok {
		return game, nil
	}

	game, err := mongo_getGame(id)
	if err != nil {
		return nil, err
	}

	registerGame(game)
	return game, nil
}

// Submits a guess to a game.
// Updates the game's Guesses and State fields.
func makeGuess(g *structs.Game, guess string) error {
	if len(guess) != g.Metadata.WordLength {
		return fmt.Errorf(`guess "%s" is not of length %d`, guess, g.Metadata.WordLength)
	}
	if !g.IsOngoing() {
		return fmt.Errorf(`game has already finished with state "%s"`, g.State)
	}

	g.Guesses = append(g.Guesses, [2]string{guess, getCorrections(guess, g.Word)})
	updateGameState(g)

	return nil
}

// Returns the corrections needed when given a guess and a correct string of the same length.
func getCorrections(guess string, correct string) string {
	guessRunes := []rune(guess)
	correctRunes := []rune(correct)

	correctFrequencies := make(map[rune]int)
	for _, char := range correctRunes {
		correctFrequencies[char] = correctFrequencies[char] + 1
	}

	corrections := make([]rune, len(guess))
	for i, ch := range guessRunes {
		if guessRunes[i] == correctRunes[i] {
			corrections[i] = 'G'
			correctFrequencies[ch] -= 1
		}
	}
	for i, ch := range guessRunes {
		if guessRunes[i] != correctRunes[i] {
			if correctFrequencies[ch] > 0 {
				corrections[i] = 'Y'
				correctFrequencies[ch] -= 1
			} else {
				corrections[i] = 'X'
			}
		}
	}

	return string(corrections)
}

// Updates the State field of a game by checking the most recent guess.
// Options: "won", "lost", and "ongoing" (no change).
func updateGameState(g *structs.Game) {
	prevCorrections := g.Guesses[len(g.Guesses)-1][1]

	switch {
	case strings.Count(prevCorrections, "G") == len(prevCorrections):
		g.State = "won"
	case len(g.Guesses) == g.Metadata.MaxGuesses:
		g.State = "lost"
	}
}
