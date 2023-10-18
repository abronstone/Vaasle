package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

/*
Word structure schema
*/
type Word struct {
	Word     string `bson:"word"`
	Length   int    `bson:"length"`
	Language string `bson:"language"`
}

func getWords(c *gin.Context) {
	/*
		Gets all of the words of a certain length from Mongo

		@param: word length (int)
		@return: list of Word objects (json[])
	*/

	// Get word length parameter
	wordLength, err := strconv.Atoi(c.Param("length"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"please enter a valid word length, error: ": err.Error()})
		return
	}

	// Get database and words collection
	database := client.Database("VaasDatabase")
	word_collection := database.Collection("words")

	// Run "find" query on words collection
	cursor, err := word_collection.Find(context.TODO(), bson.M{"length": wordLength})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve words from mongo: " + err.Error()})
		return
	}

	defer cursor.Close(context.Background())

	// Append results to output list
	var words []interface{}
	for cursor.Next(context.TODO()) {
		var word Word
		if err := cursor.Decode(&word); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decode word from mongo: " + err.Error()})
			return
		}
		words = append(words, word)
	}

	// Return words
	c.JSON(http.StatusOK, words)
}

func insertWord(c *gin.Context) {
	/*
		Inserts the inputted word into the MongoDB collection based on word length

		@param: word to be inserted (string)
		@return: confirmation message (string)
	*/

	// Get database
	database := client.Database("VaasDatabase")

	// Get correct collection
	word_parameter := c.Param("word")
	word_collection := database.Collection("words")

	// Create item
	item := Word{Word: word_parameter, Length: len(word_parameter)}

	// Insert item
	word_collection.InsertOne(context.TODO(), item)

	// Return confirmation message
	c.JSON(http.StatusOK, map[string]string{"message": "Word inserted successfully"})
}
