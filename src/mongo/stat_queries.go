package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func mostCommonWords(c *gin.Context) {
	/*
		Gets the top 5 most commonly used words by a player with the following aggregation pipeline on the games collection:

			1. Match games from a specific user ID
			2. Add new field for number of guesses
			3. Map first element of each guess to a new array
			4. Unwind arrays of guesses to unique documents
			5. Group documents together with count included
			6. Add 'word' field
			7. Sort result by count
			8. Limit result to 5 documents

		@param: user ID in path parameter
		@return: array of strings with top 5 words
	*/

	// Get metadata from HTTP body
	database := client.Database("VaasDatabase")
	userID := c.Param("userid")

	gameCollection := database.Collection("games")
	cursor, err := gameCollection.Aggregate(context.TODO(), bson.A{
		// Stage 1
		bson.D{{"$match", bson.D{{"metadata.userId", userID}}}},
		// Stage 2
		bson.D{{"$addFields", bson.D{{"numGuesses", bson.D{{"$size", "$guesses"}}}}}}, bson.D{{"$match", bson.D{
			{"guesses",
				bson.D{{"$gte", bson.A{"$numGuesses", 0}}},
			},
		}}},
		// Stage 3
		bson.D{
			{"$project",
				bson.D{
					{"_id", 0},
					{"firstGuesses",
						bson.D{
							{"$map",
								bson.D{{"input", "$guesses"}, {"as", "guess"}, {"in", bson.D{{"$arrayElemAt", bson.A{"$$guess", 0}}}}},
							},
						},
					},
				},
			},
		},
		// Stage 4
		bson.D{
			{"$unwind",
				bson.D{
					{"path", "$firstGuesses"},
					{"includeArrayIndex", "string"},
					{"preserveNullAndEmptyArrays", true},
				},
			},
		},
		// Stage 5
		bson.D{
			{"$group",
				bson.D{
					{"_id", "$firstGuesses"},
					{"count", bson.D{{"$sum", 1}}},
				},
			},
		},
		// Stage 6
		bson.D{{"$addFields", bson.D{{"word", "$_id"}}}},
		// Stage 7
		bson.D{{"$sort", bson.D{{"count", 1}}}},
		// Stage 8
		bson.D{{"$limit", 5}},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "aggregation error: " + err.Error()})
		return
	}

	// Decode cursor results to array
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cursor decoding error: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}
