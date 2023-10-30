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
	router.PUT("/create-user/:username", createUser)
	router.PUT("/login/:username", logIn)

	router.Run("0.0.0.0:80")
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"message": "User container working properly"})
}

func createUser(c *gin.Context) {
	/*
		Validates if the username exists in the database by querying the Mongo service. If it does not, a new user is created and sent to the Mongo service. If the user exists, it is not sent to Mongo.

		@param: username (string)
		@return:
			- http status 200 if the credentials pass validation and user is created successfully
			- http status 401 if the credentials do not pass validation
			- http status 500 if some other problem occurred
	*/
	username := c.Param("username")

	existingUserEndpoint := "http://mongo:8000/get-user/" + username

	// Call the mongo service to retrieve the user if it exists
	res, err := http.Get(existingUserEndpoint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("No call to exsiting user endpoint" + err.Error())
		return
	}

	if res != nil {
		// Only close the body if res is not nil
		defer res.Body.Close()
	}

	// If mongo returns status code 404, call the mongo service to insert a new user and return status code 200. Otherwise, a user was found by mongo, so return status code 401.
	if res.StatusCode == http.StatusNotFound {
		newUserEndpoint := "http://mongo:8000/new-user/" + username

		// Create new user
		new_user := structs.User{
			UserName:     username,
			Games:        []string{},
			NumGames:     0,
			TotalGuesses: 0,
			Playing:      false,
		}
		userJson, err := json.Marshal(new_user)
		if err != nil {
			fmt.Println("Error marshaling new user to json" + err.Error())
			c.JSON(http.StatusInternalServerError, err)
		}

		// Create request
		byteBuffer := bytes.NewBuffer(userJson)
		request, err := http.NewRequest(http.MethodPut, newUserEndpoint, byteBuffer)
		if err != nil {
			fmt.Println("Error making mongo request" + err.Error())
			c.JSON(http.StatusInternalServerError, err)
		}
		request.Header.Set("Content-Type", "application/json")

		// Send request
		client := &http.Client{}
		_, err = client.Do(request)
		if err != nil {
			fmt.Println("Error from mongo request" + err.Error())
			c.JSON(http.StatusInternalServerError, err)
		}

		// Return status code 200 if new user was inserted successfully
		c.JSON(http.StatusOK, structs.Message{Message: "Account created successfully"})
	} else {
		// Return status code 400 if user was already found in database
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
