package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"vaas/structs"
)

var engineURL = "http://engine:5001"

func engine_newGame(metadata structs.GameMetadata) (structs.Game, error) {
	/*
		This function takes in a metadata structure, and calls Engine to create a new game to return a new Game struct

		@param: metadata for new game (structs.Metadata)
		@return: new game created by engine (structs.Game)
	*/
	newGame := structs.Game{}
	bodyBytes, err := json.Marshal(metadata)
	if err != nil {
		return newGame, err
	}

	bodyBuffer := bytes.NewBuffer(bodyBytes)

	res, err := http.Post(engineURL+"/newGame", "application/json", bodyBuffer)

	if err != nil {
		return newGame, err
	}

	defer res.Body.Close()

	responseBodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		return newGame, err
	}

	err = json.Unmarshal(responseBodyBytes, &newGame)
	if err != nil {
		return newGame, err
	}

	return newGame, nil
}

func engine_getGame(id string) (*structs.Game, error) {
	/*
		This function sends a request to Engine with the game id parameter to receive the game structure associated with that game id

		@param: game id (string)
		@return: game returned by engine (*structs.Game)
	*/
	endpoint := engineURL + "/getGame/" + id

	res, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

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
