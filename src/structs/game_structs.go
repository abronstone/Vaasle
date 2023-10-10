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
