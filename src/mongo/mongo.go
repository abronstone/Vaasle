package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Game struct {
	Id      int
	Element string
}

func getDatabase() (*mongo.Client, error) {
	// MongoDB connection string
	CONNECTION_STRING := "mongodb://vaas1:pass@localhost:27017/"

	// Set up client options
	clientOptions := options.Client().ApplyURI(CONNECTION_STRING)

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

func main() {
	// Get database
	client, err := getDatabase()
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("vaasdb")

	db.CreateCollection(context.TODO(), "test_go_collection")
	cur_collection := db.Collection("test_go_collection")
	item1 := Game{Id: 1, Element: "First Go Element"}
	item2 := Game{Id: 2, Element: "Second Go Element"}
	docs := []interface{}{
		item1,
		item2,
	}
	result2, err := cur_collection.InsertMany(context.TODO(), docs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted item: %v\n", result2)

}
