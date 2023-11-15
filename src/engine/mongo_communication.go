package main

import (
	"errors"
	"vaas/structs"
)

// Submits a new game to the Mongo API (mongo.go).
//
// The API takes a metadata struct and returns an initialized game.
// The API also initializes an empty game with this information in MongoDB.
func mongo_submitNewGame(metadata structs.GameMetadata) (string, error) {

	game := structs.Game{}
	// Make PUT request, encoding 'metadata' to request body, decode response body into 'game' of type 'structs.Game'
	_, err := structs.MakePutRequest[structs.Game]("http://mongo:8000/new-game/", metadata, &game)
	if err != nil {
		return "", err
	}
	// defer res.Body.Close()
	// game := structs.Game{}
	// err = structs.DecodeResponseBody(res, &game)
	// if err != nil {
	// 	return "", err
	// }

	if len(game.Word) == 0 {
		return "", errors.New("could not retrieve word from database")
	}

	return game.Word, nil
}

// Asks the Mongo API (mongo.go) for the game stored under the given ID.
func mongo_getGame(id string) (*structs.Game, error) {
	game := structs.Game{}
	// Make GET request, decode response into 'game' of type 'structs.Game'
	_, err := structs.MakeGetRequest[structs.Game]("http://mongo:8000/get-game/"+id, &game)
	if err != nil {
		return nil, err
	}

	return &game, nil
}

// Updates the Mongo API (mongo.go) with the new state of the given game.
func mongo_updateGame(game *structs.Game) error {
	result := structs.Message{}
	// Make PUT request, encoding 'game' to request body, decode response body into 'result' of type 'structs.Message'
	_, err := structs.MakePutRequest[structs.Message]("http://mongo:8000/update-game/", &game, &result)
	if err != nil {
		return err
	}

	return nil
}

// Asks the Mongo API (mongo.go) to make changes to the state of the given user.
func mongo_updateUser(userId string, userUpdate *structs.UserUpdate) error {

	result := structs.Message{}
	// Make POST request, encoding 'userUpdate' to request body, decode response body into 'result' of type 'structs.Message'
	res, err := structs.MakePostRequest[structs.Message]("http://mongo:8000/update-user/"+userId, userUpdate, &result)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if result.Message != "user updated successfully" {
		return errors.New("failed to send user updates to Mongo API")
	}

	return nil
}

func mongo_verifyWord(word string) (bool, error) {
	result := structs.Message{}
	// Make GET request, decode response into 'result' of type 'structs.Message'
	_, err := structs.MakeGetRequest[structs.Message]("http://mongo:8000/check-if-valid-word/"+word+"/", &result)
	if err != nil {
		return false, err
	}
	if result.Message != "The word exists in the word collection" {
		return false, nil
	}

	return true, nil
}
