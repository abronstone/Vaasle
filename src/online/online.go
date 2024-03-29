/*
	online.go
	Vaasle 2023
*/

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	/*
		Defines router endpoints
	*/
	router := gin.Default()
	router.GET("/", home)

	// User Methods
	router.PUT("/create-user", createUser)
	router.PUT("/login/:username", logIn)

	// Multiplayer Methods
	router.POST("/newMultiplayerGame/", newMultiplayerGame)
	router.GET("/getMultiplayerGame/:id", getMultiplayerGame)
	router.PUT("/joinMultiplayerGame/:id", joinMultiplayerGame)
	router.PUT("/startMultiplayerGame/:id", startMultiplayerGame)
	router.POST("/refreshMultiplayerGame/:id", refreshMultiplayerGame)

	router.Run("0.0.0.0:8000")
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"message": "User container working properly"})
}
