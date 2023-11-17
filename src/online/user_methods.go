package main

import (
	"fmt"
	"net/http"
	"vaas/structs"

	"github.com/gin-gonic/gin"
)

func createUser(c *gin.Context) {
	// Bind incoming JSON to struct
	var requestBody structs.NewUserRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Make GET request, no decoding needed
	res, err := structs.MakeGetRequest[any]("http://mongo:8000/get-user/"+requestBody.Id, nil)
	if err != nil && res == nil {
		fmt.Println("No call to existing user endpoint: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Close the response body
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		// User not found, create a new one
		newUser := structs.User{
			Id:               requestBody.Id,
			UserName:         requestBody.UserName,
			NumGamesStarted:  0,
			NumGamesFinished: 0,
			NumGamesWon:      0,
			NumGamesLost:     0,
			TotalGuesses:     0,
		}
		// Make PUT request, encoding 'newUser' to request body, no response body decoding needed
		res, err := structs.MakePutRequest[any]("http://mongo:8000/new-user/", newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer res.Body.Close()

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
	user := structs.User{}
	// Make GET request, decode response into 'user' of type 'structs.User'
	res, err := structs.MakeGetRequest[structs.User]("http://mongo:8000/get-user/"+username, &user)
	if res != nil && res.StatusCode == http.StatusNotFound {
		c.JSON(http.StatusUnauthorized, structs.Message{Message: "Login unsuccessful"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	defer res.Body.Close()

	// If the user does not exist respond with status code 401. Otherwise, respond with status code 200

	c.JSON(http.StatusOK, structs.Message{Message: "Login successful"})
}
