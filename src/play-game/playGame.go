package main

import (
	"io"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func pingEngine(c *gin.Context) {
	res, err := http.Get("http://engine:5001/")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Ping to engine failed"})
		return
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Ping to engine failed"})
		return
	}

	c.JSON(http.StatusOK, string(bodyBytes))
}

func home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, map[string]string{"message": "Play game is running"})
}

func main() {
	router := gin.Default()

	// Add CORS middleware
	router.Use(cors.Default())

	router.GET("/", home)
	router.GET("/pingEngine", pingEngine)

	router.Run("0.0.0.0:5001")
}
