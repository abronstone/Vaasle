package structs

import (
	"time"
)

// The metadata of a Wordle game.
type GameMetadata struct {
	GameID      string    `json:"gameID" bson:"gameid"`
	DateCreated time.Time `json:"dateCreated" bson:"datecreated"`
	WordLength  int       `json:"wordLength" bson:"wordlength"`
	MaxGuesses  int       `json:"maxGuesses" bson:"maxguesses"`
}

// A Wordle game.
type Game struct {
	Metadata GameMetadata `json:"metadata" bson:"metadata"`
	Guesses  [][2]string  `json:"guesses" bson:"guesses"`
	State    string       `json:"state" bson:"state"`
	Word     string       `json:"word" bson:"word"`
}

// A guess in a Wordle game.
type Guess struct {
	Id    string `json:"id"`
	Guess string `json:"guess"`
}

// A simple message.
type Message struct {
	Message string `json:"message"`
}

// Obfuscate the word of a Game.
func (g *Game) ObfuscateWord() *Game {
	return &Game{
		Metadata: g.Metadata,
		Guesses:  g.Guesses,
		State:    g.State,
		Word:     "",
	}
}
