package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// Submits a new game to the Mongo API (mongo.go).
//
// The API takes the given ID and wordLength and returns a word for this game.
// The API also initializes an empty game with this information in MongoDB.
func submitNewGame(id string, wordLength int) (string, error) {
	return "PIZZA", nil // temporary default return, as the endpoint below is not yet implemented

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
