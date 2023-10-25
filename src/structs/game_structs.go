package structs

import (
	"time"
)

// The metadata of a Wordle game.
type GameMetadata struct {
	GameID      string    `json:"gameID" bson:"gameid"`
	DateCreated time.Time `json:"dateCreated" bson:"datecreated"`
	UserName    string    `json:"username" bson:"username"`
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

type User struct {
	UserName     string   `json:"username" bson:"username"`
	Password     string   `json:"password" bson:"password"`
	Games        []string `json:"games" bson:"games"`
	NumGames     int      `json:"numgames" bson:"numgames"`
	TotalGuesses int      `json:"totalguesses" bson:"totalguesses"`
	Playing      bool     `json:"playing" bson:"playing"`
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
