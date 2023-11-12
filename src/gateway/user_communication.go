package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	req, err := http.NewRequest(http.MethodPut, "http://online:80/create-user", bytes.NewBuffer(jsonData))

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

func api_login(c *gin.Context) {
	var userName string = c.Param("username")

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new request
	req, err := http.NewRequest(http.MethodPut, "http://online:80/login/"+userName, nil)

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
