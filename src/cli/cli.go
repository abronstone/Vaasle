package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"vaas/structs"
)

func main() {
	err := ping_play_game()

	if err != nil {
		return
	}

	var currentGame structs.Game = initialize_new_game()

	fmt.Println("New game created with ID:", currentGame.ID)

	gameID := currentGame.ID

	// Make guesses until the game is won or lost (make_guess returns nil)
	for {
		fmt.Println("Guess a word:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		guess := scanner.Text()
		var cleanedUpGuess = strings.ReplaceAll(guess, " ", "")

		if len(cleanedUpGuess) != len(currentGame.Metadata.WordLength) {
			fmt.Println("Your word doesn't match the required length. Try again.")
			continue
		}

		// Make a guess and get the corrections
		var lastGuess structs.Guess = make_guess(gameID, guess)

		if lastGuess == nil {
			return
		}

		// Print out the corrections
		fmt.Println("Corrections:")
		for _, correction := range lastGuess.Corrections {
			switch correction {
			case "G":
				fmt.Println("🟩")
			case "Y":
				fmt.Println("🟧")
			default:
				fmt.Println("⬛")
			}
		}

		// Print out the current state of the game
		fmt.Println("Current state:", currentGame.State)

		// Print out the current guesses of the game
		fmt.Println("Current guesses:")
		for _, guess := range currentGame.Guesses {
			fmt.Println(guess)
		}
	}

}

func ping_play_game() error {
	fmt.Println("Sending GET request to play-game...")
	res, err := http.Get("http://play-game:5001/")
	if err != nil {
		fmt.Println("The GET request to play-game threw an error:", err)
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("The GET request to play-game threw an error:", err)
		return err
	} else {
		fmt.Println("The GET request to play-game returned:", string(body))
	}

	return nil
}

func initialize_new_game() (*structs.Game, error) {
	res, err := http.Post("http://play:5001/newGame", "application/json", nil)

	// If play-game is down, return an error
	if err != nil {
		fmt.Println("Failed to create a new game")
		return
	}

	defer res.Body.Close()

	// Create a newGame variable and unmarshal the response body into it
	newGame := structs.Game{}

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("error: Failed read response body from play-game")
		return
	}

	// format the json response data from bodybytes, have it conform to Game type. Make it filled.
	err = json.Unmarshal(bodyBytes, &newGame)
	if err != nil {
		fmt.Println("error: Failed to unmarshal response body from engine")
		return
	}

	return &newGame, nil
}

func make_guess(gameID string, guess string) (*structs.Guess, error) {
	// Create a new guess struct
	newGuess := structs.Guess{
		GameID: gameID,
		Guess:  guess,
	}

	res, err := http.Post("http://play:5001/makeGuess", "application/json", newGuess)
	// If play-game is down, return an error
	if err != nil {
		fmt.Println("Failed to create a new game")
		return
	}

	defer res.Body.Close()

	// Create a newGame variable and unmarshal the response body into it
	currentGame := structs.Game{}

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("error: Failed read response body from play-game")
		return
	}

	// Format the json response data from bodybytes, have it conform to Game type. Make it filled.
	err = json.Unmarshal(bodyBytes, &currentGame)
	if err != nil {
		fmt.Println("error: Failed to unmarshal response body from play-game")
		return
	}

	// Check if the game is won or lost
	if currentGame.State == "won" {
		fmt.Print("Congratulations! You won the game!")
		return nil
	}
	if currentGame.State == "lost" {
		fmt.Print("Unlucky! You suck at the game delete wordle and your mother hates you!")
		return nil
	}

	// Get the corrections for hinting of the last guess made
	var lastGuess structs.Guess = currentGame.Guesses[len(currentGame.Guesses)-1][1]

	return lastGuess, nil
}

// func main() {
// 	var words []string = []string{"happy", "crane", "sugar", "towel", "gator", "apple"}
// 	randomIdx := rand.Intn(len(words))

// 	randomWord := words[randomIdx]

// 	fmt.Printf("here is my randowm word  %s \n", randomWord)

// 	scanner := bufio.NewScanner(os.Stdin)

// 	var combinedInput string
// 	maxAttempts := 6

// 	for attempts := 0; attempts < maxAttempts; attempts++ {
// 		for {
// 			fmt.Printf("Please enter a word that has %d letters: ", len(randomWord))
// 			scanner.Scan()
// 			input := scanner.Text()

// 			// Remove all spaces
// 			combinedInput = strings.ReplaceAll(input, " ", "")

// 			if len(combinedInput) == len(randomWord) {
// 				break
// 			} else {
// 				fmt.Println("Your word doesn't match the required length. Try again.")
// 			}
// 		}
// 		score := 0

// 		output := ""

// 		for i := 0; i < len(randomWord); i++ {
// 			if combinedInput[i] == randomWord[i] {
// 				output += "🟩"
// 				score += 2

// 			} else if strings.Contains(randomWord, string(combinedInput[i])) {
// 				output += "🟧"
// 				score += 1
// 			} else {
// 				output += "⬛"
// 			}
// 		}

// 		fmt.Println("Result:", output)
// 		fmt.Println("Score:", score)

// 		if combinedInput == randomWord {
// 			fmt.Println("congratulations! you guessed the word correctly!")
// 			break
// 		} else {
// 			fmt.Println("game has ended, better luck next time")

//			}
//		}
//	}
