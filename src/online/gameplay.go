package main

import (
	"net/http"
	"vaas/structs"

	"github.com/gin-gonic/gin"
)

func startMultiplayerGame(c *gin.Context) {
	c.JSON(http.StatusOK, structs.Message{Message: "startMultiplayerGame working"})
}

func refreshMultiplayerGame(c *gin.Context) {
	c.JSON(http.StatusOK, structs.Message{Message: "refreshMultiplayerGame working"})
}
