package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"vaas/structs"
)

func mongo_createMultiplayerGame(multiplayerGame structs.MultiplayerGame) error {
	/*
		Takes in a multiplayer game struct and sends it to Mongo container to insert into database

		@param: multiplayer game struct (structs.MultiplayerGame)
		@return: success response (string)
	*/

	// Marshal multiplayerGame
	bodyBytes, err := json.Marshal(multiplayerGame)
	if err != nil {
		return err
	}
	multiplayerGameBodyBuffer := bytes.NewBuffer(bodyBytes)

	// Call Mongo's "newMultiplayerGame" endpoint
	endpoint := "http://mongo:8000/intializeMultiplayerGame/"
	req, err := http.NewRequest(http.MethodPut, endpoint, multiplayerGameBodyBuffer)
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

func mongo_getMultiplayerGame(multiplayerGameID string) (*structs.MultiplayerGame, error) {
	/*
		Takes in a multiplayer game ID and returns the multiplayer game associated with it from Mongo

		@param: multiplayer game id (string)
		@return: multiplayer game (structs.MultiplayerGame)
	*/
	endpoint := "http://mongo:8000/getMultiplayerGame/" + multiplayerGameID

	res, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	multiplayerGame := &structs.MultiplayerGame{}
	err = json.Unmarshal(bodyBytes, multiplayerGame)
	if err != nil {
		return nil, err
	}

	return multiplayerGame, nil
}

func mongo_addUserToMultiplayerGame(multiplayerGameID string, game structs.Game) error {
	/*
		Takes in a multiplayer game ID, a new individual game ID, and a user ID to send to Mongo. Mongo should add the game ID and user ID to the 'games' map in the multiplayer game struct associated with the multiplayer game ID, and returns a response based off Mongo's response

		@param: multiplayer game id (string), game id (string), user id (string)
		@return: success response (string)
	*/
	endpoint := "http://mongo:8000/registerUserInMultiplayerGame/" + multiplayerGameID

	bodyBytes, err := json.Marshal(game)
	if err != nil {
		return err
	}
	gameBodyBuffer := bytes.NewBuffer(bodyBytes)

	req, err := http.NewRequest(http.MethodPut, endpoint, gameBodyBuffer)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("Could not create multiplayer game due to Mongo error")
	}

	return nil
}
