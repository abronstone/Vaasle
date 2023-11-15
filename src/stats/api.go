package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"vaas/structs"

	"github.com/gin-gonic/gin"
)

/*
The following code is a result of the collaboration of Team Vaas and
Peter Kelly from team Lelkolopher (Dominion)
*/

// Stats API's main method.
func main() {
	router := gin.Default()

	router.GET("/", api_home)
	router.GET("/getStats/:userId", api_getStats)
	router.GET("/getLeaderboard/:maxNumUsers", api_getLeaderboard)

	router.Run("0.0.0.0:5001")
}

// Default home endpoint for checking if the stats container is running.
func api_home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "stats is running"})
}

// Gets relevant stats for the given user.
func api_getStats(c *gin.Context) {
	userId := c.Param("userId")

	// TODO: query the mongo container to get relevant stats for user: userId
	// TODO: aggregate data into presentable statistics for the frontend to display

	// Asks the Mongo API (mongo.go) for the game stored under the given ID.
	// func mongo_getGame(id string) (*structs.Game, error) {
	// 1. Send request

	endpoint := "http://mongo:8000/get-user/" + userId

	res, err := http.Get(endpoint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: "Error getting user stats" + err.Error()})
	}
	defer res.Body.Close()

	// 2. Parse response body
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: "Error getting user stats" + err.Error()})
	}

	user := &structs.User{}
	err = json.Unmarshal(bodyBytes, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: "Error getting user stats" + err.Error()})
	}

	games := user.Games
	gamesPlayed := user.NumGames
	gamesWon := 0
	guesses := map[string]int{}

	for element := range games {
		var thisgame *structs.Game
		thisgame, err = mongo_getGame(element)
		if err != nil {
			c.JSON(http.StatusInternalServerError, structs.Message{Message: "Error getting user stats" + err.Error()})
		}
		for _, value := range thisgame.Guesses {
			count, ok := guesses[value[0]]
			if ok {
				guesses[value[0]] = count + 1
			}
			if !ok {
				guesses[value[0]] = 1
			}

		}
		if thisgame.State == "won" {
			gamesWon += 1
		}
	}

	var winPercentage = float32(gamesWon) / float32(gamesPlayed)

	var mostCommonWord = ""
	var num = 0
	for key, value := range guesses {
		if value > num {
			mostCommonWord = key
			num = value
		}
	}
	returnval := structs.IndividualUserStats{GamesPlayed: gamesPlayed, WinPercentage: winPercentage, MostGuessedWord: mostCommonWord}
	// returnval.GamesPlayed = gamesPlayed
	// returnval.WinPercentage = winPercentage
	// returnval.MostGuessedWord = mostCommonWord

	c.JSON(http.StatusOK, returnval)
}

// Asks the Mongo API (mongo.go) for the game stored under the given ID.
func mongo_getGame(id string) (*structs.Game, error) {
	// 1. Send request
	endpoint := "http://mongo:8000/get-game/" + id

	res, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// 2. Parse response body
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	game := &structs.Game{}
	err = json.Unmarshal(bodyBytes, game)
	if err != nil {
		return nil, err
	}

	return game, nil
}

// Returns the top maxNumUsers in descending order with respect to total games played.
func api_getLeaderboard(c *gin.Context) {
	maxNumUsers, err := strconv.Atoi(c.Param("maxNumUsers"))
	if err != nil {
		errorMsg := fmt.Sprintf("error: maxNumUsers must be an int, got %s", c.Param("maxNumUsers"))
		c.JSON(http.StatusBadRequest, structs.Message{Message: errorMsg})
		return
	}

	// TODO: delete this
	message := fmt.Sprintf("received request for leaderboard of %d users", maxNumUsers)
	c.JSON(http.StatusOK, structs.Message{Message: message})

	// TODO: query the mongo container to get the top maxNumUsers with respect to total games played.
}
