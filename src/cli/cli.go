package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Guess struct {
	Id    string `json:"id"`
	Guess string `json:"guess"`
}

type Game struct {
	Metadata GameMetadata `json:"metadata" bson:"metadata"`
	Guesses  [][2]string  `json:"guesses" bson:"guesses"`
	State    string       `json:"state" bson:"state"`
	Word     string       `json:"word" bson:"word"`
}

type GameMetadata struct {
	GameID     string `json:"gameID" bson:"gameid"`
	WordLength int    `json:"wordLength" bson:"wordlength"`
	MaxGuesses int    `json:"maxGuesses" bson:"maxguesses"`
}

func main() {
	err := ping_play_game()

	if err != nil {
		fmt.Println("Failed to ping gateway")
		return
	}

	var wordLength int
	var maxGuesses int

	// Allow user to choose word length and max guesses, within certain bounds
	scanner := bufio.NewScanner(os.Stdin)
	for wordLength < 5 || wordLength > 6 {
		fmt.Println("Please enter a word length of 5 or 6:")
		scanner.Scan()
		var err error
		wordLength, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
		}
	}

	for maxGuesses < 1 || maxGuesses > 10 {
		fmt.Println("Please enter a max number of guesses between 1 and 10:")
		scanner.Scan()
		var err error
		maxGuesses, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
		}
	}

	currentGame, err := initialize_new_game(wordLength, maxGuesses)

	if currentGame == nil {
		return
	}

	fmt.Println("New game created with ID:", currentGame.Metadata.GameID)

	gameID := currentGame.Metadata.GameID

	// Make guesses until the game is won or lost (make_guess returns nil)
	for {
		fmt.Println("Guess a word:")
		scanner.Scan()
		guess := scanner.Text()
		var cleanedUpGuess = strings.ReplaceAll(guess, " ", "")

		if len(cleanedUpGuess) != currentGame.Metadata.WordLength {
			fmt.Println("Your word doesn't match the required length of", currentGame.Metadata.WordLength, "Try again.")
			continue
		}

		// Make a guess and get the corrections
		lastGuess, status, err := make_guess(gameID, guess)

		if err != nil {
			return
		}

		output := ""
		// Print out the corrections
		for _, correction := range lastGuess {
			switch string(correction) {
			case "G":
				output += "🟩"
			case "Y":
				output += "🟧"
			default:
				output += "⬛"
			}
		}

		fmt.Println(output)

		if status == "won" || status == "lost" {
			return
		}
	}

}

func ping_play_game() error {
	fmt.Println("Sending GET request to gateway...")
	res, err := http.Get("http://gateway:5001/")
	if err != nil {
		fmt.Println("The GET request to gateway threw an error:", err)
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("The GET request to gateway threw an error:", err)
		return err
	} else {
		fmt.Println("The GET request to gateway returned:", string(body))
	}

	return nil
}

func initialize_new_game(wordLength int, maxGuesses int) (*Game, error) {
	// Word length can only be 5 or 6 b/c those are the only sized words we have
	// in the DB at the moment

	resPayload := GameMetadata{
		WordLength: wordLength,
		MaxGuesses: maxGuesses,
	}

	// Convert the payload to JSON
	payloadBytes, err := json.Marshal(resPayload)
	if err != nil {
		fmt.Println("Failed to build a request body to create a new game")
		return nil, err
	}

	// Create a buffer with the JSON data
	bodyBuffer := bytes.NewBuffer(payloadBytes)

	// Make request to gateway to create a new game with a request body
	res, err := http.Post("http://gateway:5001/newGame", "application/json", bodyBuffer)

	// If gateway is down, return an error
	if err != nil {
		fmt.Println("Failed to create a new game")
		return nil, err
	}

	defer res.Body.Close()

	// Create a newGame variable and unmarshal the response body into it
	newGame := Game{}

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("error: Failed read response body from gateway")
		return nil, err
	}

	// format the json response data from bodybytes, have it conform to Game type. Make it filled.
	err = json.Unmarshal(bodyBytes, &newGame)
	if err != nil {
		fmt.Println("error: Failed to unmarshal response body from engine")
		return nil, err
	}

	return &newGame, nil
}

func make_guess(gameID string, guess string) (string, string, error) {
	// Make a guess of type Guess
	guessStruct := Guess{
		Id:    gameID,
		Guess: guess,
	}
	// Convert guessStruct to JSON
	jsonData, err := json.Marshal(guessStruct)
	if err != nil {
		fmt.Println("Failed to marshal guess:", err)
		return "", "error", err
	}

	// Use strings.NewReader to convert the JSON string to an io.Reader
	res, err := http.Post("http://gateway:5001/makeGuess", "application/json", strings.NewReader(string(jsonData)))

	if err != nil {
		fmt.Println("Failed to create a new game")
		return "", "error", err
	}

	defer res.Body.Close()

	// Create a newGame variable and unmarshal the response body into it
	currentGame := Game{}

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("error: Failed read response body from gateway")
		return "", "error", err
	}

	// Format the json response data from body bytes, have it conform to Game type. Make it filled.
	err = json.Unmarshal(bodyBytes, &currentGame)
	if err != nil {
		fmt.Println("error: Failed to unmarshal response body from gateway")
		return "", "error", err
	}

	// Get the corrections for hinting of the last guess made
	var lastGuessCorrections string = currentGame.Guesses[len(currentGame.Guesses)-1][1]

	// Check if the game is won or lost
	if currentGame.State == "won" {
		fmt.Println("Congratulations! You won the game!")
		return lastGuessCorrections, "won", nil
	}
	if currentGame.State == "lost" {
		fmt.Println("Unlucky! Try again!")
		fmt.Println("The word was:", currentGame.Word)
		return lastGuessCorrections, "lost", nil
	}

	return lastGuessCorrections, "ongoing", nil
}
