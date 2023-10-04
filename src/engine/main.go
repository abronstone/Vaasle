package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, map[string]string{"message": "Engine is running"})
}

func main() {
	router := gin.Default()

	// Add CORS middleware
	router.Use(cors.Default())

	router.GET("/", home)
	router.Run("0.0.0.0:5000")
}
