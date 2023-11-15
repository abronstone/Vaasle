package main

import (
	"encoding/json"
	"io"
	"net/http"
	"vaas/structs"
)

func mongo_getUser(userId string) (*structs.User, error) {
	endpoint := "http://mongo:8000/get-user/" + userId

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

	user := &structs.User{}
	err = json.Unmarshal(bodyBytes, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Asks the Mongo API (mongo.go) for the most common words for the given user.
func mongo_getCommonWords(userId string) ([]structs.CommonWordFrequency, error) {
	// 1. Send request
	endpoint := "http://mongo:8000/most-common-words/" + userId

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

	words := make([]structs.CommonWordFrequency, 0)
	err = json.Unmarshal(bodyBytes, &words)
	if err != nil {
		return nil, err
	}

	return words, nil
}
