package main

import (
	"net/http"
	"vaas/structs"

	"github.com/gin-gonic/gin"
)

func startSharedGame(c *gin.Context) {
	c.JSON(http.StatusOK, structs.Message{Message: "startSharedGame working"})
}

func refreshSharedGame(c *gin.Context) {
	c.JSON(http.StatusOK, structs.Message{Message: "refreshSharedGame working"})
}
