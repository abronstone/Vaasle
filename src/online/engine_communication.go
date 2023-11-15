package main

import (
	"vaas/structs"
)

var engineURL = "http://engine:5001"

func engine_newGame(metadata structs.GameMetadata) (*structs.Game, error) {
	/*
		This function takes in a metadata structure, and calls Engine to create a new game to return a new Game struct

		@param: metadata for new game (structs.Metadata)
		@return: new game created by engine (structs.Game)
	*/
	newGame := structs.Game{}
	// Make POST request, encoding 'metadata' to request body, decode response body into 'newGame' of type 'structs.Game'
	_, err := structs.MakePostRequest[structs.Game]("http://engine:5001/newGame", metadata, &newGame)
	if err != nil {
		return nil, err
	}

	return &newGame, nil
}

func engine_getGame(id string) (*structs.Game, error) {
	game := structs.Game{}
	// Make GET request, decode response into 'game' of type 'structs.Game'
	_, err := structs.MakeGetRequest[structs.Game](engineURL+"/getGame/"+id, &game)
	if err != nil {
		return nil, err
	}

	return &game, nil
}
