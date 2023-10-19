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
	Word     string       `json:"-" bson:"word"`
}

// A Wordle game with an exportable public Word.
// Only for Mongo and Engine communication.
type GameExposed struct {
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
	Games        []string `json:"games" bson:"games"`
	NumGames     int      `json:"numgames" bson:"numgames"`
	TotalGuesses int      `json:"totalguesses" bson:"totalguesses"`
}

// Convert a GameExposed to a Game (remove the ability to export Word to JSON).
func (g *GameExposed) ConvertToGame() *Game {
	return &Game{
		Metadata: g.Metadata,
		Guesses:  g.Guesses,
		State:    g.State,
		Word:     g.Word,
	}
}

// Convert a Game to a GameExposed (add the ability to export Word to JSON).
func (g *Game) ConvertToGameExposed() *GameExposed {
	return &GameExposed{
		Metadata: g.Metadata,
		Guesses:  g.Guesses,
		State:    g.State,
		Word:     g.Word,
	}
}
