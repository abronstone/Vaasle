package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"vaas/structs"

	"github.com/gin-gonic/gin"
)

func api_createMultiplayerGame(c *gin.Context) {
	newGameCustomMetadata := structs.GameMetadata{}
	if err := c.ShouldBindJSON(&newGameCustomMetadata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request format, expected {wordLength: int, maxGuesses: int, userID: string}"})
		return
	}

	bodyBytes, err := json.Marshal(newGameCustomMetadata)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal the request body: " + err.Error()})
		return
	}
	bodyBuffer := bytes.NewBuffer(bodyBytes)

	res, err := http.Post("http://online:8000/newMultiplayerGame", "application/json", bodyBuffer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}
	defer res.Body.Close()

	game := structs.MultiplayerGame{}

	responseBodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}

	err = json.Unmarshal(responseBodyBytes, &game)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, game.GetShareable())
}

func api_joinMultiplayerGame(c *gin.Context) {
	id := c.Param("id")

	newGameCustomMetadata := structs.GameMetadata{}
	if err := c.ShouldBindJSON(&newGameCustomMetadata); err != nil {
		c.JSON(http.StatusBadRequest, structs.Message{Message: "Bad request format, expected {wordLength: int, maxGuesses: int, userID: string}"})
		return
	}

	// Create a new HTTP client
	client := &http.Client{}
	endpoint := "http://online:8000/joinMultiplayerGame/" + id

	// Convert struct to JSON
	jsonData, err := json.Marshal(newGameCustomMetadata)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}

	// Create a new request
	req, err := http.NewRequest(http.MethodPut, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}

	// Set Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Execute the request
	res, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: "could not join game"})
		return
	}

	responseBodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}

	game := &structs.MultiplayerGame{}
	err = json.Unmarshal(responseBodyBytes, game)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, game.GetShareable())
}

func api_startMultiplayerGame(c *gin.Context) {
	id := c.Param("id")
	client := &http.Client{}
	endpoint := "http://online:8000/startMultiplayerGame/" + id

	// Create a new request
	req, err := http.NewRequest(http.MethodPut, endpoint, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}

	// Set Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Execute the request
	res, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: "could not join game"})
		return
	}
}

func api_refreshMultiplayerGame(c *gin.Context) {
	c.JSON(http.StatusOK, structs.Message{Message: "refresh endpoint called"})
}
