package main

import (
	"encoding/json"
	"io"
	"net/http"
	"vaas/structs"
)

func mongo_getUser(userid string) (*structs.User, error) {
	/*
		Takes in a userid string and requests the user struct from Mongo associated with that user id

		@param: user id (string)
		@return: user (structs.User)
	*/
	res, err := http.Get("http://mongo:8000/get-user/" + userid)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	user := structs.User{}
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func mongo_createMultiplayerGame(multiplayerGame structs.MultiplayerGame) error {
	/*
		Takes in a multiplayer game struct and sends it to Mongo container to insert into database

		@param: multiplayer game struct (structs.MultiplayerGame)
		@return: success response (string)
	*/

	// Make PUT request, encoding 'multiplayerGame' to request body, no response body decoding needed
	res, err := structs.MakePutRequest[any]("http://mongo:8000/initializeMultiplayerGame/", &multiplayerGame)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func mongo_getMultiplayerGame(multiplayerGameID string) (*structs.MultiplayerGame, error) {
	/*
		Takes in a multiplayer game ID and returns the multiplayer game associated with it from Mongo

		@param: multiplayer game id (string)
		@return: multiplayer game (structs.MultiplayerGame)
	*/
	multiplayerGame := structs.MultiplayerGame{}
	// Make GET request, decode response into 'multiplayerGame' of type 'structs.MultiplayerGame'
	_, err := structs.MakeGetRequest[structs.MultiplayerGame]("http://mongo:8000/getMultiplayerGame/"+multiplayerGameID, &multiplayerGame)
	if err != nil {
		return nil, err
	}

	return &multiplayerGame, nil
}

func mongo_addUserToMultiplayerGame(multiplayerGameID string, game *structs.Game) error {
	/*
		Takes in a multiplayer game ID, a new individual game ID, and a user ID to send to Mongo. Mongo should add the game ID and user ID to the 'games' map in the multiplayer game struct associated with the multiplayer game ID, and returns a response based off Mongo's response

		@param: multiplayer game id (string), game id (string), user id (string)
		@return: success response (string)
	*/
	// Make PUT request, encoding 'game' to request body, no response body decoding needed
	res, err := structs.MakePutRequest[any]("http://mongo:8000/registerUserInMultiplayerGame/"+multiplayerGameID, game)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}

func mongo_startMultiplayerGame(multiplayerGameID string) error {
	/*
		Takes in a multiplayer game ID, and requests Mongo to start that game.
		This entails setting the state of the game to "ongoing".

		@param: multiplayer game id (string)
		@return: error or nil if successful
	*/

	// Make PUT request, no encoding/decoding needed
	res, err := structs.MakePutRequest[any]("http://mongo:8000/startMultiplayerGame/"+multiplayerGameID, nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func mongo_updateMultiplayerGame(multiplayerGameID string, update *structs.MultiplayerGameUpdate) error {

	// Make PUT request, encoding 'update' to request body, no response body decoding needed
	res, err := structs.MakePutRequest[any]("http://mongo:8000/updateMultiplayerGame/"+multiplayerGameID, update)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
