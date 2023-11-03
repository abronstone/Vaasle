package main

import (
	"fmt"
	"net/http"
	"strconv"
	"vaas/structs"

	"github.com/gin-gonic/gin"
)

// Stats API's main method.
func main() {
	router := gin.Default()

	router.GET("/", api_home)
	router.GET("/getStats/:userId", api_getStats)
	router.GET("/getLeaderboard/:maxNumUsers", api_getLeaderboard)

	router.Run("0.0.0.0:5001")
}

// Default home endpoint for checking if the stats container is running.
func api_home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "stats is running"})
}

// Gets relevant stats for the given user.
func api_getStats(c *gin.Context) {
	userId := c.Param("userId")

	// TODO: delete this
	message := fmt.Sprintf("%s won π games", userId)
	c.JSON(http.StatusOK, structs.Message{Message: message})

	// TODO: query the mongo container to get relevant stats for user: userId
	// TODO: aggregate data into presentable statistics for the frontend to display
}

// Returns the top maxNumUsers in descending order with respect to total games played.
func api_getLeaderboard(c *gin.Context) {
	maxNumUsers, err := strconv.Atoi(c.Param("maxNumUsers"))
	if err != nil {
		errorMsg := fmt.Sprintf("error: maxNumUsers must be an int, got %s", c.Param("maxNumUsers"))
		c.JSON(http.StatusBadRequest, structs.Message{Message: errorMsg})
		return
	}

	// TODO: delete this
	message := fmt.Sprintf("received request for leaderboard of %d users", maxNumUsers)
	c.JSON(http.StatusOK, structs.Message{Message: message})

	// TODO: query the mongo container to get the top maxNumUsers with respect to total games played.
}
