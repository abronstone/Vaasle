package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"vaas/structs"
)

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

	res, err := http.Post("http://engine:5001/newGame", "application/json", bodyBuffer)

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
