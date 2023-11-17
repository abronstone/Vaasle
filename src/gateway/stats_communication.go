package main

import (
	"encoding/json"
	"io"
	"net/http"
	"vaas/structs"

	"github.com/gin-gonic/gin"
)

/*
Gets stats for an individual user

@param: user id in path parameter
@return: user statistics in HTTP response (structs.IndividualUserStats)
*/
func api_getStats(c *gin.Context) {
	stats, err := getStats(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, stats)
}

/*
Requests the stats container with user id, returns individual status struct

@param: user id (string)
@return: user statistics (structs.IndividualUserStats)
*/
func getStats(userId string) (*structs.IndividualUserStats, error) {
	// 1. Send request
	endpoint := "http://stats:5001/getStats/" + userId

	res, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// 2. Parse response body
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	stats := &structs.IndividualUserStats{}
	err = json.Unmarshal(bodyBytes, stats)
	if err != nil {
		return nil, err
	}

	return stats, nil
}
