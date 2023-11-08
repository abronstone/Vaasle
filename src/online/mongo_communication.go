package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"vaas/structs"
)

func mongo_createSharedGame(sharedGame structs.SharedGame) error {
	/*
		Takes in a shared game struct and sends it to Mongo container to insert into database

		@param: shared game struct (structs.SharedGame)
		@return: success response (string)
	*/

	// Marshal sharedGame
	bodyBytes, err := json.Marshal(sharedGame)
	if err != nil {
		return err
	}
	sharedGameBodyBuffer := bytes.NewBuffer(bodyBytes)

	// Call Mongo's "newSharedGame" endpoint
	endpoint := "http://mongo:8000/newSharedGame/"
	req, err := http.NewRequest(http.MethodPut, endpoint, sharedGameBodyBuffer)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Return response based on Mongo response
	if res.StatusCode != http.StatusOK {
		return errors.New("Could not create share game due to Mongo error")
	}

	return nil
}

func mongo_getSharedGame(sharedGameID string) (*structs.SharedGame, error) {
	/*
		Takes in a shared game ID and returns the shared game associated with it from Mongo

		@param: shared game id (string)
		@return: shared game (structs.SharedGame)
	*/
	endpoint := "http://mongo:8000/getSharedGame/" + sharedGameID

	res, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	sharedGame := &structs.SharedGame{}
	err = json.Unmarshal(bodyBytes, sharedGame)
	if err != nil {
		return nil, err
	}

	return sharedGame, nil
}

func mongo_addUserToSharedGame(sharedGameID string, gameID string, userID string) error {
	/*
		Takes in a shared game ID, a new individual game ID, and a user ID to send to Mongo. Mongo should add the game ID and user ID to the 'games' map in the shared game struct associated with the shared game ID, and returns a response based off Mongo's response

		@param: shared game id (string), game id (string), user id (string)
		@return: success response (string)
	*/
	endpoint := "http://mongo:8000/addUserToSharedGame/" + sharedGameID + "/" + gameID + "/" + userID

	req, err := http.NewRequest(http.MethodPost, endpoint, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("Could not create share game due to Mongo error")
	}

	return nil
}
