package main

import (
	"bytes"
	"encoding/json"
	"net/http"

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

	// Create a new HTTP client
	client := &http.Client{}

	// Convert struct to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error converting data to JSON"})
		return
	}

	// Create a new request
	req, err := http.NewRequest(http.MethodPut, "http://online:8000/create-user", bytes.NewBuffer(jsonData))

	// Handle request creation error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a new request: " + err.Error()})
		return
	}

	// Set Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Execute the request
	res, err := client.Do(req)

	// Handle execution error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute user creation request: " + err.Error()})
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

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new request
	req, err := http.NewRequest(http.MethodPut, "http://online:8000/login/"+userName, nil)

	// Handle request creation error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a new request: " + err.Error()})
		return
	}

	// Execute the request
	res, err := client.Do(req)

	// Handle execution error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute login request: " + err.Error()})
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
