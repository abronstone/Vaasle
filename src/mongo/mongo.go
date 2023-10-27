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
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
		return nil
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client
}

// Initialize global Mongo client
var client *mongo.Client = getDatabase()

func initializeDB(c *gin.Context) {
	/*
		Clears all collections in database (in future will just clear game history)

		@param: nothing
		@return: confirmation message
	*/
	db := client.Database("VaasDatabase")
	// deleteMany function without a filter deletes all documents in a collection
	_, err := db.Collection("games").DeleteMany(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to clear games collection: " + err.Error()})
		return
	}
	// _, err = db.Collection("words").DeleteMany(context.Background(), bson.M{})
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to clear words collection: " + err.Error()})
	// 	return
	// }
	c.JSON(http.StatusOK, map[string]string{"message": "RESET DATABASE"})
}

func home(c *gin.Context) {
	/*
		Help message to display routes
	*/
	c.JSON(http.StatusOK, map[string]string{
		"GET: /initialize-db":      "CLEARS DATABASE COLLELCTIONS (use with caution)",
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
	router.GET("/check-if-valid-word/:word", checkIfValidWord)

	router.PUT("/new-game/", newGame)
	router.GET("/get-game/:id", getGame)
	router.PUT("/update-game/", updateGameState)

	router.PUT("/new-user/:username", newUser)
	router.GET("/get-user/:username", getUser)
	router.POST("/update-user/:username", updateUser)

	router.Run("0.0.0.0:8000")
}
