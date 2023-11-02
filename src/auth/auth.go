package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"vaas/structs"

	"github.com/gin-gonic/gin"
)

func main() {
	/*
		Defines router endpoints
	*/
	router := gin.Default()
	router.GET("/", home)
	router.PUT("/create-user", createUser)
	router.PUT("/login/:username", logIn)

	router.Run("0.0.0.0:80")
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"message": "User container working properly"})
}

func createUser(c *gin.Context) {
	// UserRequestBody holds incoming request body data
	type UserRequestBody struct {
		UserName string `json:"userName"`
		Id       string `json:"id"`
	}

	// Bind incoming JSON to struct
	var requestBody UserRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Use the id from the request body as the user ID

	existingUserEndpoint := "http://mongo:8000/get-user/" + requestBody.Id

	// Check for existing user
	res, err := http.Get(existingUserEndpoint)
	if err != nil {
		fmt.Println("No call to existing user endpoint: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Close the response body
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		// User not found, create a new one
		newUserEndpoint := "http://mongo:8000/new-user/"

		newUser := structs.User{
			Id:           requestBody.Id,
			UserName:     requestBody.UserName,
			Games:        []string{},
			NumGames:     0,
			TotalGuesses: 0,
			Playing:      false,
		}

		userJson, err := json.Marshal(newUser)
		if err != nil {
			fmt.Println("Error marshaling new user to JSON: " + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Create and send request to create new user
		request, err := http.NewRequest(http.MethodPut, newUserEndpoint, bytes.NewBuffer(userJson))
		if err != nil {
			fmt.Println("Error making request to Mongo: " + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		_, err = client.Do(request)
		if err != nil {
			fmt.Println("Error from Mongo request: " + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, structs.Message{Message: "Account created successfully"})
	} else {
		// User already exists
		c.JSON(http.StatusBadRequest, structs.Message{Message: "Error: Account already exists"})
	}
}

func logIn(c *gin.Context) {
	/*
		Validates if the username exists in the database by querying the Mongo service. If the user does not exist a 401 error is thrown.

		@param: username (string)
		@return:
			- http status 200 if the credentials are authenticated via validation
			- http status 401 if the credentials are not authenticated via validation
			- http status 500 if some other problem occurred
	*/
	username := c.Param("username")

	// Call the mongo service to retrieve the user if it exists
	existingUserEndpoint := "http://mongo:8000/get-user/" + username
	res, err := http.Get(existingUserEndpoint)
	if err != nil || res == nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	user := &structs.User{}
	err = json.Unmarshal(bodyBytes, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	// If the user does not exist respond with status code 401. Otherwise, respond with status code 200
	if res.StatusCode == http.StatusNotFound {
		c.JSON(http.StatusUnauthorized, structs.Message{Message: "Login unsuccessful"})
	}
}
