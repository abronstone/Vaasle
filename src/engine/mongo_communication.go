package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

// Submits a new game to the Mongo API (mongo.go).
//
// The API takes the given ID and wordLength and returns a word for this game.
// The API also initializes an empty game with this information in MongoDB.
func mongo_submitNewGame(id string, wordLength int) (string, error) {
	return string(make([]byte, wordLength)), nil // temporary default return, as the endpoint below is not yet implemented

	res, err := http.Get("http://mongo:5000/newGame/" + id + "/" + strconv.Itoa(wordLength))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	word := struct {
		Word string `json:"word"`
	}{}
	err = json.Unmarshal(bodyBytes, &word)
	if err != nil {
		return "", err
	}

	return word.Word, nil
}

// Asks the Mongo API (mongo.go) for the game stored under the given ID.
func mongo_getGame(id string) (*game, error) {
	res, err := http.Get("http://mongo:5000/getGame/" + id)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	game := game{}
	err = json.Unmarshal(bodyBytes, &game)
	if err != nil {
		return nil, err
	}

	return &game, nil
}

// Updates the Mongo API (mongo.go) with the new state of the given game.
func mongo_updateGame(game *game) error {
	return nil // temporary default return, as the endpoint below is not yet implemented

	gameJson, err := json.Marshal(game)
	if err != nil {
		return err
	}

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, "http://mongo:5000/getGame/", bytes.NewBuffer(gameJson))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	result := struct {
		Message string `json:"message"`
	}{}
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return err
	}

	if result.Message != "success" {
		return errors.New("failed to send game updates to Mongo API")
	}

	return nil
}
