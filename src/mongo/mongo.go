/*
	mongo.go
	Last updated 10/11/2023

	To run:
		go mod download
		go run mongo.go

	Needs to be ran in src/mongo in vaas-final project

	vaas.ai 2023
	Advanced Software Design: Fall 2023
	Carleton College
*/

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
	Word     string `bson:"word"`
	Length   int    `bson:"length"`
	Language string `bson:"language"`
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
	item := Word{Word: word_parameter, Length: len(word_parameter)}

	// Insert item
	word_collection.InsertOne(context.TODO(), item)

	// Return confirmation message
	c.JSON(http.StatusOK, map[string]string{"message": "Word inserted successfully"})
}

func newGame(c *gin.Context) {
	/*
		Creates a new game structure with the specified parameters, gets a random word that conforms to the word length (soon to be language), and returns the word

		OVERWRITE SAME GAME IF IT EXISTS

		@param: game metadata structure in HTTP body
		@return: word
	*/

	// Get metadata from HTTP body
	headerVal := c.GetHeader("Content-Type")
	fmt.Println(headerVal)
	var metadata structs.GameMetadata
	if err := c.ShouldBindJSON(&metadata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gameID := metadata.GameID
	wordLength := metadata.WordLength

	// Get c ollections
	database := client.Database("VaasDatabase")
	wordCollection := database.Collection("words")
	gameCollection := database.Collection("games")

	deleteFilter := bson.D{{"metadata.gameid", bson.D{{"$eq", gameID}}}}
	gameCollection.DeleteOne(context.TODO(), deleteFilter)

	// Get a random word
	var randomWord bson.M
	cursor, err := wordCollection.Aggregate(context.TODO(), bson.A{
		bson.D{{"$match", bson.D{{"length", wordLength}}}},
		bson.D{{"$sample", bson.D{{"size", 1}}}},
	})
	if err != nil {
		log.Fatal(err)
	}
	if cursor.Next(context.TODO()) {
		if err := cursor.Decode(&randomWord); err != nil {
			log.Fatal(err)
		}
	}
	defer cursor.Close(context.Background())

	// Create new game structure and insert into database
	gameMetadata := structs.GameMetadata{GameID: gameID, WordLength: wordLength, MaxGuesses: 6}
	guesses := [][2]string{}
	game := structs.GameExposed{Word: randomWord["word"].(string), Metadata: gameMetadata, Guesses: guesses, State: "ongoing"}
	gameCollection.InsertOne(context.TODO(), game)

	// Return random word
	c.JSON(http.StatusOK, game)
}

func updateGameState(c *gin.Context) {
	/*
		Updates the game state and guesses

		@param: updated game structure in HTTP body
		@return: JSON confirmation message
	*/

	// Gets the HTTP header and body
	headerVal := c.GetHeader("Content-Type")
	fmt.Println(headerVal)
	var gameData structs.Game
	if err := c.ShouldBindJSON(&gameData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newState := gameData.State
	newGuesses := gameData.Guesses
	gameID := gameData.Metadata.GameID

	// Update document in database
	database := client.Database("VaasDatabase")
	gameCollection := database.Collection("games")
	filter := bson.D{{"metadata.gameid", gameID}}
	update := bson.D{{"$set", bson.D{{"state", newState}, {"guesses", newGuesses}}}}

	_, err := gameCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	// Return confirmation message
	c.JSON(http.StatusOK, bson.M{"message": "Game updated successfully"})
}

func getGame(c *gin.Context) {
	/*
		Finds and returns the specified game from the database

		@param: game ID
		@return: game structure with the associated game ID
	*/

	// Get game ID from path parameters
	gameID := c.Param("id")

	database := client.Database("VaasDatabase")
	gameCollection := database.Collection("games")

	// Create match pipeline stage
	matchStage := bson.A{
		bson.D{{"$match", bson.D{{"metadata.gameid", gameID}}}},
	}
	var game structs.GameExposed
	// Run aggregation
	cursor, err := gameCollection.Aggregate(context.TODO(), matchStage)
	defer cursor.Close(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	if cursor.Next(context.TODO()) {
		if err := cursor.Decode(&game); err != nil {
			log.Fatal(err)
		}
	}

	// Return game
	c.JSON(http.StatusOK, game)
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
	cursor, err := word_collection.Find(context.TODO(), bson.M{"length": wordLength})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	// Append results to output list
	var words []interface{}
	for cursor.Next(context.TODO()) {
		var word Word
		if err := cursor.Decode(&word); err != nil {
			log.Fatal(err)
		}
		words = append(words, word)
	}

	// Return words
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
	/*
		Help message to display routes
	*/
	c.JSON(http.StatusOK, map[string]string{
		"GET: /initialize-db":      "CLEARS DATABASE COLELCTIONS (use with caution)",
		"GET: /insert-word/<word>": "Inserts word into database",
		"GET: /get-words/<length>": "Gets words of parameter length",
		"PUT: /new-game/":          "Creates a new game based on HTTP body metadata",
		"GET: /get-game/<id>":      "Returns the game with the parameter ID",
		"PUT: /update-game/":       "Updates the game state and guesses based on HTTP body"})
}

func main() {
	/*
		GIN Router mapping
	*/
	router := gin.Default()
	router.GET("/", home)
	router.GET("/initialize-db", initializeDB)

	router.GET("/get-words/:length", getWords)
	router.GET("/insert-word/:word", insertWord)

	router.PUT("/new-game/", newGame)
	router.GET("/get-game/:id", getGame)
	router.PUT("/update-game/", updateGameState)

	router.Run("0.0.0.0:8000")
}
