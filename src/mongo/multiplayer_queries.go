package main

import (
	"context"
	"net/http"
	"vaas/structs"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// PUT: Called by the Online container.
// Initializes a new multiplayer game when a "creating" user creates one.
// Expects a MultiplayerGame struct in the request body.
func api_initializeMultiplayerGame(c *gin.Context) {
	game := &structs.MultiplayerGame{}
	if err := c.ShouldBindJSON(game); err != nil {
		c.JSON(http.StatusBadRequest, structs.Message{Message: err.Error()})
		return
	}
	if err := mongodb_insertMultiplayerGame(game); err != nil {
		c.JSON(http.StatusBadRequest, structs.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, structs.Message{Message: "multiplayer game initialized"})
}

// PUT: Called by the Online container.
// Changes an existing multiplayer game's state to "ongoing".
func api_startMultiplayerGame(c *gin.Context) {
	filter := bson.M{"multiplayergameid": c.Param("id")}
	update := bson.M{"$set": bson.M{"state": "ongoing"}}

	if err := mongodb_updateMultiplayerGame(&filter, &update); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, structs.Message{Message: "multiplayer game started"})
}

// GET: Called by the Online container.
// Returns an unobfuscated representation of a multiplayer game.
func api_getMultiplayerGame(c *gin.Context) {
	game, err := mongodb_getMultiplayerGame(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, game)
}

// PUT: Called by the Online container.
// Registers a "joining" user's gameId into a pre-existing Multiplayer game.
// Expects a Game struct in the request body.
func api_registerUserInMultiplayerGame(c *gin.Context) {
	game := &structs.Game{}
	if err := c.ShouldBindJSON(game); err != nil {
		c.JSON(http.StatusBadRequest, structs.Message{Message: err.Error()})
		return
	}

	filter := bson.M{"multiplayergameid": c.Param("id")}
	mapKey := "games" + "." + game.Metadata.UserId
	update := bson.M{"$set": bson.M{mapKey: game.Metadata.GameID}}

	if err := mongodb_updateMultiplayerGame(&filter, &update); err != nil {
		c.JSON(http.StatusBadRequest, structs.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, structs.Message{Message: "added user to multiplayer game"})
}

// PUT: Called by the Online container.
// Updates a multiplayer game's state/winner but does nothing if already finished.
// Expects a MultiplayerGameUpdate struct in the request body.
func api_updateMultiplayerGame(c *gin.Context) {
	gameUpdate := &structs.MultiplayerGameUpdate{}
	if err := c.ShouldBindJSON(gameUpdate); err != nil {
		c.JSON(http.StatusBadRequest, structs.Message{Message: err.Error()})
		return
	}

	filter := bson.M{"multiplayergameid": c.Param("id"), "state": "ongoing"}
	update := bson.M{"$set": gameUpdate}

	if err := mongodb_updateMultiplayerGame(&filter, &update); err != nil {
		c.JSON(http.StatusBadRequest, structs.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, structs.Message{Message: "updated state of multiplayer game"})
}

// Helper function to get a multiplayer game.
func mongodb_getMultiplayerGame(id string) (*structs.MultiplayerGame, error) {
	collection := client.Database("VaasDatabase").Collection("multiplayergames")

	filter := bson.M{"multiplayergameid": id}
	game := &structs.MultiplayerGame{}

	err := collection.FindOne(context.TODO(), filter).Decode(&game)
	if err != nil {
		return nil, err
	}
	return game, nil
}

// Helper function to insert a multiplayer game.
func mongodb_insertMultiplayerGame(game *structs.MultiplayerGame) error {
	collection := client.Database("VaasDatabase").Collection("multiplayergames")
	_, err := collection.InsertOne(context.TODO(), game)
	return err
}

// Helper function to update a multiplayer game.
func mongodb_updateMultiplayerGame(filter *bson.M, update *bson.M) error {
	collection := client.Database("VaasDatabase").Collection("multiplayergames")
	_, err := collection.UpdateOne(context.TODO(), *filter, *update)
	return err
}
