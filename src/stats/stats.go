package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getStats(c *gin.Context) {

}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "Stats container is running",
	})
}

func main() {
	router := gin.Default()

	router.GET("/", home)
	router.GET("/get-stats/:userid", getStats)

	router.Run("0.0.0.0:8001")
}
