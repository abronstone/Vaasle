package structs

// The metadata of a Wordle game.
type GameMetadata struct {
	GameID     string `json:"gameID" bson:"gameid"`
	WordLength int    `json:"wordLength" bson:"wordlength"`
	MaxGuesses int    `json:"maxGuesses" bson:"maxguesses"`
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

// Convert a GameExposed to a Game (remove the ability to export Word to JSON).
func (g *GameExposed) ConvertToGame() *Game {
	return &Game{
		Metadata: g.Metadata,
		Guesses:  g.Guesses,
		State:    g.State,
		Word:     g.Word,
	}
}
