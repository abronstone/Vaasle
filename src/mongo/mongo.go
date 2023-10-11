package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"vaas/structs"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
Word structure schema
*/
type Word struct {
	Word       string `bson:"word"`
	WordLength int    `bson:"wordLength"`
}

func getDatabase() *mongo.Client {
	/*
		Returns a MongoDB Client instance
	*/
	// MongoDB connection string
	CONNECTION_STRING := "mongodb+srv://vaas_admin:adv1software2design3@vaasdatabase.sarpr4r.mongodb.net"

	// Set up client options
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(CONNECTION_STRING).SetServerAPIOptions(serverAPI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client
}

// Initialize global Mongo client
var client *mongo.Client = getDatabase()

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
	item := Word{Word: word_parameter, WordLength: len(word_parameter)}

	// Insert item
	word_collection.InsertOne(context.TODO(), item)

	c.JSON(http.StatusOK, map[string]string{"message": "Word inserted successfully"})
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
		log.Fatal(err)
	}

	// Get database and words collection
	database := client.Database("VaasDatabase")
	word_collection := database.Collection("words")

	// Run "find" query on words collection
	cursor, err := word_collection.Find(context.TODO(), bson.M{"wordlength": wordLength})
	if err != nil {
		log.Fatal(err)
	}

	// Append results to output list
	var words []interface{}
	for cursor.Next(context.TODO()) {
		var word Word
		if err := cursor.Decode(&word); err != nil {
			log.Fatal(err)
		}
		words = append(words, word)
	}

	c.JSON(http.StatusOK, words)
}

func initializeDB(c *gin.Context) {
	/*
		Clears all collections in database (in future will just clear game history)

		@param: nothing
		@return: confirmation message
	*/
	db := client.Database("VaasDatabase")
	// deleteMany function without a filter deletes all documents in a collection
	_, err := db.Collection("words").DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, map[string]string{"message": "RESET DATABASE"})
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"/": "This message", "/initialize-db": "CLEARS DATABASE COLELCTIONS (use with caution)", "/insert-word/<word>": "Inserts word into database", "/get-words/<length>": "Gets words of parameter length"})
}

func main() {
	/*
		GIN Router mapping
	*/
	router := gin.Default()

	router.GET("/", home)
	router.GET("/initialize-db", initializeDB)
	router.GET("/insert-word/:word", insertWord)
	router.GET("/get-words/:length", getWords)

	router.Run("0.0.0.0:5000")

	// Proof of concept, please remove
	game := structs.Game{}
	log.Println(len(game.Guesses))
}
