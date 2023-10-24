package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Stats API's main method.
func main() {
	router := gin.Default()

	router.GET("/", api_home)
	router.GET("/getStats/:username", api_getStats)

	router.Run("0.0.0.0:5001")
}

// Default home endpoint for checking if the stats container is running.
func api_home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "stats is running"})
}

// Gets relevant stats for the given user.
func api_getStats(c *gin.Context) {
	// Want to make a request to Mongo API to get user's information

	// Then, filter this information and reorganize it into a helpful format for statistics.

	c.JSON(http.StatusOK, gin.H{"message": "user won π games"})
}
