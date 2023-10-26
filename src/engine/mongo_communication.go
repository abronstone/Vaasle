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
	// 1. Prepare request headers and body
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

	// 2. Send request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// 3. Parse response body
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	game := structs.Game{}
	err = json.Unmarshal(bodyBytes, &game)
	if err != nil {
		return "", err
	}

	if len(game.Word) == 0 {
		return "", errors.New("could not retrieve word from database")
	}

	return game.Word, nil
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

// Updates the Mongo API (mongo.go) with the new state of the given game.
func mongo_updateGame(game *structs.Game) error {
	// 1. Prepare request headers and body
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

	// 2. Send request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// 3. Parse response body
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

// Asks the Mongo API (mongo.go) to make changes to the state of the given user.
func mongo_updateUser(username string, userUpdate *structs.UserUpdate) error {
	// 1. Send request
	endpoint := "http://mongo:8000/update-game/" + username

	bodyBytes, err := json.Marshal(userUpdate)
	if err != nil {
		return err
	}

	bodyBuffer := bytes.NewBuffer(bodyBytes)

	res, err := http.Post(endpoint, "application/json", bodyBuffer)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	// 2. Parse response body
	bodyBytes, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	result := &structs.Message{}
	err = json.Unmarshal(bodyBytes, result)
	if err != nil {
		return err
	}

	if result.Message != "user updated successfully" {
		return errors.New("failed to send user updates to Mongo API")
	}

	return nil
}
