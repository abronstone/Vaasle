package main

import (
	"net/http"
	"vaas/structs"

	"github.com/gin-gonic/gin"
)

/*
Creates a new user in the system.
Receives a request with user details, validates the data, and sends a PUT request to the user creation endpoint.
Returns a success message upon successful creation, or an error message in case of failure or if the user already exists.

@param: User details in the request body (username and ID)
@return: Success or error message
*/
func api_newUser(c *gin.Context) {

	// Define a struct to bind the request body
	type NewUserRequestBody struct {
		UserName string `json:"userName"`
		Id       string `json:"id"`
	}

	var requestBody NewUserRequestBody

	// Bind JSON to struct
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Make PUT request, encode 'requestBody' into request, no decode needed
	res, err := structs.MakePutRequest[structs.Game]("http://online:8000/create-user", requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "PUT request error"})
		return
	}

	// Close the response body
	defer res.Body.Close()

	// Check for status codes
	if res.StatusCode == http.StatusBadRequest {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	if res.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

/*
Handles user login.
Extracts the username from the URL parameter, sends a PUT request to the login endpoint.
Returns a success message if login is successful, or an error message if the user does not exist or if there are other login issues.

@param: Username from the URL parameter
@return: Login success or error message
*/
func api_login(c *gin.Context) {
	var userName string = c.Param("username")

	// Make PUT request, no encoding/decoding needed
	res, err := structs.MakePutRequest[any]("http://online:8000/login/"+userName, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "PUT request error while logging in"})
		return
	}

	// Close the response body
	defer res.Body.Close()

	// Check for status codes
	if res.StatusCode == http.StatusUnauthorized {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
		return
	}

	if res.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error logging in"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
