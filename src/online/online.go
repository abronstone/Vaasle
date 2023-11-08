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

	router.PUT("/newSharedGame/", newSharedGame)
	router.GET("/getSharedGame/:id", getSharedGame)
	router.POST("/joinSharedGame", joinSharedGame)

	router.PUT("/startSharedGame/:id", startSharedGame)

	router.POST("/refreshSharedGame", refreshSharedGame)

	router.Run("0.0.0.0:8000")
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"message": "User container working properly"})
}
