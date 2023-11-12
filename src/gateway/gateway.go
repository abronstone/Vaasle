package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Add CORS middleware
	router.Use(cors.Default())

	router.GET("/", api_home)

	// Engine
	router.GET("/pingEngine", api_pingEngine)
	router.POST("/newGame", api_newGame)
	router.GET("/getGame/:id", api_getGame)
	router.POST("/makeGuess", api_makeGuess)

	// Users
	router.PUT("/createUser", api_newUser)
	router.PUT("/login/:username", api_login)

	// Online
	router.POST("/createMultiplayerGame", api_createMultiplayerGame)
	router.PUT("/joinMultiplayerGame/:id", api_joinMultiplayerGame)
	router.PUT("/startMultiplayerGame/:id", api_startMultiplayerGame)
	router.POST("/refreshMultiplayerGame/:id", api_refreshMultiplayerGame)

	router.Run("0.0.0.0:5001")
}

func api_home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, map[string]string{"message": "Gateway is running"})
}
