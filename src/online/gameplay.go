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

func refreshMultiplayerGame(c *gin.Context) {
	id := c.Param("id")
	multiplayerGame, err := helper_getMultiplayerGame(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}

	games := make(map[string]*structs.Game)
	populateGames(multiplayerGame, games)
	userNames, err := getUserNames(multiplayerGame.Games)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: "Error mapping user ids to usernames: " + err.Error()})
		return
	}

	// multiplayer game was already discovered to be finished by someone else
	if multiplayerGame.State == "won" || multiplayerGame.State == "lost" {
		c.JSON(http.StatusOK, &structs.MultiplayerFrontendUpdate{
			State:           multiplayerGame.State,
			WinnerID:        multiplayerGame.WinnerID,
			Word:            multiplayerGame.Word,
			UserCorrections: getUserCorrections(games),
			UserNames:       userNames,
		})
		return
	}

	update := getNewGameUpdate(multiplayerGame.State, games)
	if update.IsFinished() {
		if err := mongo_updateMultiplayerGame(id, update); err != nil {
			c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
			return
		}

		// Need to get game from Mongo again in case someone else's
		// online container already marked the multiplayer game as finished.
		if multiplayerGame, err = helper_getMultiplayerGame(id); err != nil {
			c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
			return
		}

		games = make(map[string]*structs.Game)
		populateGames(multiplayerGame, games)

		update = getNewGameUpdate(multiplayerGame.state, games)
	}

	word := ""
	if multiplayerGame.IsFinished() {
		word = multiplayerGame.Word
	}

	c.JSON(http.StatusOK, &structs.MultiplayerFrontendUpdate{
		State:           update.State,
		WinnerID:        update.WinnerID,
		Word:            word,
		UserCorrections: getUserCorrections(games),
		UserNames:       userNames,
	})
}

// Queries Mongo to populate the given games map using gameIDs.
func populateGames(multiplayerGame *structs.MultiplayerGame, games map[string]*structs.Game) {
	for userID, gameID := range multiplayerGame.Games {
		game, err := engine_getGame(gameID)
		// Rather than erroring out if we failed to get this user's game,
		// we will just skip this user (won't be rendered this time).
		if err == nil {
			games[userID] = game
		}
	}
}

// Analyze every user's game to generate a new multiplayer game update.
func getNewGameUpdate(state string, games map[string]*structs.Game) *structs.MultiplayerGameUpdate {
	update := structs.MultiplayerGameUpdate{
		State:    state,
		WinnerID: "",
	}
	allLost := true
	for _, game := range games {
		if game.State == "won" {
			update.State = "won"
			update.WinnerID = game.Metadata.UserId
			return &update
		}
		if game.State != "lost" {
			allLost = false
		}
	}
	if allLost {
		update.State = "lost"
	}
	return &update
}

// Generates all users' corrections from a slice of game structs.
func getUserCorrections(games map[string]*structs.Game) map[string][]string {
	userCorrections := make(map[string][]string)
	for userID, game := range games {
		userCorrections[userID] = game.GetCorrections()
	}
	return userCorrections
}

// Generate all users' usernames from a map of user ids to game ids
func getUserNames(games map[string]string) (map[string]string, error) {
	output := make(map[string]string)
	for k := range games {
		user, err := mongo_getUser(k)
		if err != nil {
			return nil, err
		}
		output[k] = user.UserName
	}
	return output, nil
}
