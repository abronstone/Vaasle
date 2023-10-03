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
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true // Allow any origin (not recommended for production)
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	router.Use(cors.New(config))

	router.GET("/", home)
	router.Run("0.0.0.0:5000")
}
