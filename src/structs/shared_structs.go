package structs

import (
	"time"
)

// The metadata of a Wordle game.
type GameMetadata struct {
	GameID       string    `json:"gameID" bson:"gameid"`
	DateCreated  time.Time `json:"dateCreated" bson:"datecreated"`
	UserId       string    `json:"userId" bson:"userId"`
	WordLength   int       `json:"wordLength" bson:"wordlength"`
	MaxGuesses   int       `json:"maxGuesses" bson:"maxguesses"`
	EnforcedWord string    `json:"enforcedWord" bson:"enforcedword"`
}

// A Wordle game.
type Game struct {
	Metadata GameMetadata `json:"metadata" bson:"metadata"`
	Guesses  [][2]string  `json:"guesses" bson:"guesses"`
	State    string       `json:"state" bson:"state"`
	Word     string       `json:"word" bson:"word"`
}

// A multiplayer game.
type MultiplayerGame struct {
	MultiplayerGameID string            `json:"multiplayerGameID" bson:"multiplayergameid"`
	HostID            string            `json:"hostID" bson:"hostid"`
	Games             map[string]string `json:"games" bson:"games"`
	State             string            `json:"state" bson:"state"`
	WinnerID          string            `json:"winnerID" bson:"winnerid"`
	Word              string            `json:"word" bson:"word"`
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
	Id               string `json:"id" bson:"id"`
	UserName         string `json:"username" bson:"username"`
	NumGamesStarted  int    `json:"numgamesstarted" bson:"numgamesstarted"`
	NumGamesFinished int    `json:"numgamesfinished" bson:"numgamesfinished"`
	NumGamesWon      int    `json:"numgameswon" bson:"numgameswon"`
	NumGamesLost     int    `json:"numgameslost" bson:"numgameslost"`
	TotalGuesses     int    `json:"totalguesses" bson:"totalguesses"`
}

// An update to make for a given user.
type UserUpdate struct {
	ChangeInNumGamesStarted  int `json:"changeInNumGamesStarted"`
	ChangeInNumGamesFinished int `json:"changeInNumGamesFinished"`
	ChangeInNumGamesWon      int `json:"changeInNumGamesWon"`
	ChangeInNumGamesLost     int `json:"changeInNumGamesLost"`
	ChangeInTotalGuesses     int `json:"changeInTotalGuesses"`
}

type NewUserRequestBody struct {
	UserName string `json:"userName"`
	Id       string `json:"id"`
}

type IndividualUserStats struct {
	GamesPlayed              int     `json:"gamesPlayed"`
	WinPercentage            float32 `json:"winPercentage"`
	MostGuessedWord          string  `json:"mostGuessedWord"`
	MostGuessedWordFrequency int     `json:"mostGuessedWordFrequency"`
}

// Obfuscate the designated word of a GameMetadata.
func (g *GameMetadata) ObfuscateWord() *GameMetadata {
	return &GameMetadata{
		GameID:       g.GameID,
		DateCreated:  g.DateCreated,
		UserId:       g.UserId,
		WordLength:   g.WordLength,
		MaxGuesses:   g.MaxGuesses,
		EnforcedWord: "",
	}
}

// Obfuscate the word of a Game.
func (g *Game) ObfuscateWord() *Game {
	return &Game{
		Metadata: *g.Metadata.ObfuscateWord(),
		Guesses:  g.Guesses,
		State:    g.State,
		Word:     "",
	}
}

// Returns whether or not a Game is ongoing.
func (g *Game) IsOngoing() bool {
	return g.State == "ongoing"
}

// Returns whether or not a MultiplayerGame has finished.
func (g *MultiplayerGame) IsFinished() bool {
	return g.State == "won" || g.State == "lost"
}

// Returns a shareable version of a Game, depending on whether it is ongoing.
func (g *Game) GetShareable() *Game {
	if g.IsOngoing() {
		return g.ObfuscateWord()
	}
	return g
}

// Returns a slice containing the corrections of a Game.
func (g *Game) GetCorrections() []string {
	corrections := make([]string, 0, len(g.Guesses))
	for _, pair := range g.Guesses {
		corrections = append(corrections, pair[1])
	}
	return corrections
}

// Makes a simple user update after a single guess.
func (g *Game) GetUserUpdateAfterGuess() *UserUpdate {
	changeInNumGamesFinished := 0
	changeInNumGamesWon := 0
	changeInNumGamesLost := 0

	// If the game is not ongoing anymore, then it has just ended.
	// This means that we can increment the number of played games by one.
	if !g.IsOngoing() {
		changeInNumGamesFinished = 1
		if g.State == "won" {
			changeInNumGamesWon = 1
		} else {
			changeInNumGamesLost = 1
		}
	}

	return &UserUpdate{
		ChangeInNumGamesFinished: changeInNumGamesFinished,
		ChangeInNumGamesWon:      changeInNumGamesWon,
		ChangeInNumGamesLost:     changeInNumGamesLost,
		ChangeInTotalGuesses:     1,
	}
}

// Obfuscate the word of a Game.
func (g *MultiplayerGame) ObfuscateWord() *MultiplayerGame {
	return &MultiplayerGame{
		MultiplayerGameID: g.MultiplayerGameID,
		HostID:            g.HostID,
		Games:             g.Games,
		State:             g.State,
		WinnerID:          g.WinnerID,
		Word:              "",
	}
}

// Returns a shareable version of a MultiplayerGame, depending on whether it is ongoing.
func (g *MultiplayerGame) GetShareable() *MultiplayerGame {
	if !g.IsFinished() {
		return g.ObfuscateWord()
	}
	return g
}

// A multiplayer game update.
type MultiplayerGameUpdate struct {
	State    string `json:"state" bson:"state"`
	WinnerID string `json:"winnerID" bson:"winnerid"`
}

// Returns whether or not a MultiplayerGameUpdate is in a finished state.
func (u *MultiplayerGameUpdate) IsFinished() bool {
	return u.State == "won" || u.State == "lost"
}

// A representation of a multiplayer game to be sent to the frontend.
type MultiplayerFrontendUpdate struct {
	State           string              `json:"state" bson:"state"`
	WinnerID        string              `json:"winnerID" bson:"winnerid"`
	Word            string              `json:"word" bson:"word"`
	UserCorrections map[string][]string `json:"userCorrections" bson:"usercorrections"`
	UserNames       map[string]string   `json:"userNames" bson:"usernames"`
}

type CommonWordFrequency struct {
	Id    string `json:"_id"`
	Count int    `json:"count"`
	Word  string `json:"word"`
}
