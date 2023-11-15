package main

import (
	"net/http"
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

	router.Run("0.0.0.0:5001")
}

// Default home endpoint for checking if the stats container is running.
func api_home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "stats is running"})
}

// Gets relevant stats for the given user.
func api_getStats(c *gin.Context) {
	user, err := mongo_getUser(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}

	words, err := mongo_getCommonWords(user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}

	mostCommonWord := ""
	highestFrequency := 0
	for _, word := range words {
		if word.Count > highestFrequency {
			mostCommonWord = word.Word
		}
	}

	winPercentage := float32(0)
	if user.NumGamesFinished != 0 {
		winPercentage = float32(user.NumGamesWon) / float32(user.NumGamesFinished)
	}

	c.JSON(http.StatusOK, &structs.IndividualUserStats{
		GamesPlayed:              user.NumGamesFinished,
		WinPercentage:            winPercentage,
		MostGuessedWord:          mostCommonWord,
		MostGuessedWordFrequency: highestFrequency,
	})
}
