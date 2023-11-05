package main

import (
	"net/http"
	"vaas/structs"

	"github.com/gin-gonic/gin"
)

func newSharedGame(c *gin.Context) {
	c.JSON(http.StatusOK, structs.Message{Message: "newSharedGame working"})
}

func getSharedGame(c *gin.Context) {
	c.JSON(http.StatusOK, structs.Message{Message: "getSharedGame working"})
}

func joinSharedGame(c *gin.Context) {
	c.JSON(http.StatusOK, structs.Message{Message: "joinSharedGame working"})
}
