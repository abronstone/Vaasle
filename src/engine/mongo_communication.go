package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"vaas/structs"
)

// Submits a new game to the Mongo API (mongo.go).
//
// The API takes a metadata struct and returns an initialized game.
// The API also initializes an empty game with this information in MongoDB.
func mongo_submitNewGame(metadata structs.GameMetadata) (string, error) {
	metadataJson, err := json.Marshal(metadata)
	if err != nil {
		return "", err
	}

	endpoint := "http://mongo:8000/new-game/"
	byteBuffer := bytes.NewBuffer(metadataJson)

	req, err := http.NewRequest(http.MethodPut, endpoint, byteBuffer)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	gameExposed := structs.GameExposed{}
	err = json.Unmarshal(bodyBytes, &gameExposed)
	if err != nil {
		return "", err
	}

	game := gameExposed.ConvertToGame()
	if len(game.Word) == 0 {
		return "", errors.New("could not retrieve word from database")
	}

	return game.Word, nil
}

// Asks the Mongo API (mongo.go) for the game stored under the given ID.
func mongo_getGame(id string) (*structs.Game, error) {
	endpoint := "http://mongo:8000/get-game/" + id

	res, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	gameExposed := structs.GameExposed{}
	err = json.Unmarshal(bodyBytes, &gameExposed)
	if err != nil {
		return nil, err
	}

	return gameExposed.ConvertToGame(), nil
}

// Updates the Mongo API (mongo.go) with the new state of the given game.
func mongo_updateGame(game *structs.Game) error {
	gameJson, err := json.Marshal(game)
	if err != nil {
		return err
	}

	endpoint := "http://mongo:8000/update-game/"
	byteBuffer := bytes.NewBuffer(gameJson)

	req, err := http.NewRequest(http.MethodPut, endpoint, byteBuffer)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	result := structs.Message{}
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return err
	}

	if result.Message != "game updated successfully" {
		return errors.New("failed to send game updates to Mongo API")
	}

	return nil
}
