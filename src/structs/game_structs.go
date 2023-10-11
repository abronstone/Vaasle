package structs

// The metadata of a Wordle game.
type GameMetadata struct {
	GameID     string `json:"gameID"`
	WordLength int    `json:"wordLength"`
	MaxGuesses int    `json:"maxGuesses"`
}

// A Wordle game.
type Game struct {
	Metadata GameMetadata `json:"metadata"`
	Guesses  [][2]string  `json:"guesses"`
	State    string       `json:"state"`
	Word     string       `json:"-"`
}

// A Wordle game with an exportable public Word.
// Only for Mongo and Engine communication.
type GameExposed struct {
	Metadata GameMetadata `json:"metadata"`
	Guesses  [][2]string  `json:"guesses"`
	State    string       `json:"state"`
	Word     string       `json:"word"`
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
