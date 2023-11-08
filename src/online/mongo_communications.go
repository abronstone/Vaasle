package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"vaas/structs"
)

func mongo_createSharedGame(sharedGame structs.SharedGame) (string, error) {
	/*
		Takes in a shared game struct and sends it to Mongo container to insert into database

		@param: shared game struct (structs.SharedGame)
		@return: success response (string)
	*/

	// Marshal sharedGame
	bodyBytes, err := json.Marshal(sharedGame)
	if err != nil {
		return "", err
	}
	sharedGameBodyBuffer := bytes.NewBuffer(bodyBytes)

	// Call Mongo's "newSharedGame" endpoint
	endpoint := "http://mongo:8000/newSharedGame/"
	req, err := http.NewRequest(http.MethodPut, endpoint, sharedGameBodyBuffer)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// Return response based on Mongo response
	if res.StatusCode == http.StatusOK {
		return "Created shared game", nil
	} else {
		return "Could not create shared game due to Mongo error", nil
	}
}
