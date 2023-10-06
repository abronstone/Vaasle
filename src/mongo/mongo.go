package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Word struct {
	WordId     int    `json:"wordID"`
	Word       string `json:"word"`
	WordLength int    `json:"wordLength"`
}

/*
Returns a MongoDB Client instance
*/
func getDatabase() (*mongo.Client, error) {
	// MongoDB connection string
	CONNECTION_STRING := "mongodb+srv://vaas_admin:adv1software2design3@vaasdatabase.sarpr4r.mongodb.net/?retryWrites=true&w=majority"

	// Set up client options
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(CONNECTION_STRING).SetServerAPIOptions(serverAPI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client, nil
}

/*
Returns a unique UUID integer
*/
func getWordId(wordIdGetter func() int) int {
	wordID := wordIdGetter()
	return wordID
}

/*
Inserts the inputted word into the MongoDB collection based on word length
*/
func insertWord(c *gin.Context) {

	naiveGetUniqueWordID := func() int {
		return int(uuid.New().ID())
	}

	// Get database
	client, err := getDatabase()
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database("VaasDatabase")

	// Get correct collection
	word_parameter := c.Param("word")
	//word_collection := database.Collection(strconv.Itoa(len(word_parameter)) + "-letter-words")
	word_collection := database.Collection("words")

	//Create item
	wordID := getWordId(naiveGetUniqueWordID)
	item := Word{WordId: wordID, Word: word_parameter, WordLength: len(word_parameter)}

	//Insert item
	word_collection.InsertOne(context.TODO(), item)
	if err != nil {
		log.Fatal(err)
	}
	//Return status
	defer client.Disconnect(context.Background())
	c.JSON(http.StatusOK, map[string]string{"message": "Word inserted successfully"})
}

func getWords(c *gin.Context) {
	wordLength, err := strconv.Atoi(c.Param("length"))
	fmt.Println("1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("1")
	// Get database
	client, err := getDatabase()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("2")
	database := client.Database("VaasDatabase")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("3")
	word_collection := database.Collection("words")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("4")
	words, err := word_collection.Distinct(context.TODO(), "word", bson.M{"wordLength": wordLength})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("5")
	defer client.Disconnect(context.Background())
	c.JSON(http.StatusOK, words)
}

/*
Clears all collections in database
*/
func initializeDB(c *gin.Context) {
	client, err := getDatabase()
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("VaasDatabase")

	// collections, err := db.ListCollectionNames(context.TODO(), bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//Clear all collections
	// for _, name := range collections {
	// 	db.Collection(name).DeleteMany(context.Background(), bson.M{})
	// }
	db.Collection("words").DeleteMany(context.Background(), bson.M{})
	defer client.Disconnect(context.Background())
	c.JSON(http.StatusOK, map[string]string{"message": "RESET DATABASE"})

}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"/": "This message", "/initialize-db": "CLEARS DATABASE COLELCTIONS (use with caution)", "/insert-word/<word>": "Inserts word into database"})
}

func main() {
	router := gin.Default()

	router.GET("/", home)
	router.GET("/initialize-db", initializeDB)
	router.GET("/insert-word/:word", insertWord)
	router.GET("/get-words/:length", getWords)

	router.Run("0.0.0.0:5000")
}
