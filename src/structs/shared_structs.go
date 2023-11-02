package structs

import (
	"time"
)

// The metadata of a Wordle game.
type GameMetadata struct {
	GameID      string    `json:"gameID" bson:"gameid"`
	DateCreated time.Time `json:"dateCreated" bson:"datecreated"`
	UserId      string    `json:"userId" bson:"userId"`
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

// A user.
type User struct {
	Id           string   `json:"id" bson:"id"`
	UserName     string   `json:"username" bson:"username"`
	Games        []string `json:"games" bson:"games"`
	NumGames     int      `json:"numgames" bson:"numgames"`
	TotalGuesses int      `json:"totalguesses" bson:"totalguesses"`
	Playing      bool     `json:"playing" bson:"playing"`
}

// An update to make for a given user.
type UserUpdate struct {
	ChangeInNumGames     int `json:"changeInNumGames"`
	ChangeInTotalGuesses int `json:"changeInTotalGuesses"`
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

// Returns whether or not a Game is ongoing.
func (g *Game) IsOngoing() bool {
	return g.State == "ongoing"
}

// Returns a shareable version of a Game, depending on whether it is ongoing.
func (g *Game) GetShareable() *Game {
	if g.IsOngoing() {
		return g.ObfuscateWord()
	}
	return g
}

// Makes a simple user update after a single guess.
func (g *Game) GetUserUpdateAfterGuess() *UserUpdate {
	changeInNumGames := 0

	// If the game is not ongoing anymore, then it has just ended.
	// This means that we can increment the number of played games by one.
	if !g.IsOngoing() {
		changeInNumGames += 1
	}

	return &UserUpdate{
		ChangeInNumGames:     changeInNumGames,
		ChangeInTotalGuesses: 1,
	}
}
