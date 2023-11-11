package main

import (
	"net/http"
	"vaas/structs"

	"github.com/gin-gonic/gin"
)

func startMultiplayerGame(c *gin.Context) {
	if err := mongo_startMultiplayerGame(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, structs.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, structs.Message{Message: "started multiplayer game"})
}

// Convert games into an array of characters for each guess
// userid: [['G','X','Y','Y','X'],['G','X','X',Y','G'],...]
func refreshMultiplayerGame(c *gin.Context) {
	multiplayerGameID := c.Param("id")
	multiplayerGame, err := mongo_getMultiplayerGame(multiplayerGameID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// var hardCodedIDS = make(map[string]string)
	// hardCodedIDS["google-oauth2|110842460204812740716"] = "f965d204-cd51-4582-91ec-4061f962adf9"
	// hardCodedIDS["yungbron"] = "f95647ae-f088-4b66-b64c-74249e591af5"
	// hardCodedIDS["google-oauth2|116861231659098811689"] = "2443f571-4907-497d-acd5-8ee812d942f4"
	guesses, allLost, winner, err := convertAndValidate(multiplayerGame.Games)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if winner != "" || allLost {
		err := mongo_updateMultiplayerGame(multiplayerGameID, allLost, winner)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	}
	state := "ongoing"
	if allLost {
		state = "lost"
	}
	if winner != "" {
		state = "won"
	}
	frontEndUpdate := structs.MultiplayerFrontendUpdate{
		State:           state,
		WinnerID:        winner,
		UserCorrections: guesses,
	}
	c.JSON(http.StatusOK, frontEndUpdate)
}

func convertAndValidate(games map[string]string) (map[string][]string, bool, string, error) {
	/*
		Converts a mapping of user ids to game ids:
			{userid: gameid}

		...to a map of user ids to array of guesses
			{userid: [[guess1],[guess2],...]}
	*/
	var winner = ""
	var allLost = true
	var guessesMap = make(map[string][]string)
	for userid, gameid := range games {
		game, err := engine_getGame(gameid)
		if err != nil {
			return nil, false, "", err
		}
		if game.State == "won" && winner == "" {
			winner = game.Metadata.UserId
		}
		if game.State != "lost" {
			allLost = false
		}
		guesses := game.Guesses
		var guessArray []string
		for _, result := range guesses {
			correction := result[1]
			// var guessArray []string
			// for _, r := range correction {
			// 	guessArray = append(guessArray, r)
			// }
			guessArray = append(guessArray, correction)
		}
		guessesMap[userid] = guessArray
	}
	return guessesMap, allLost, winner, nil
}
