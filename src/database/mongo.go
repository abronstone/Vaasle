package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Game struct {
	id      int
	element string
}

func getDatabase() (*mongo.Client, error) {
	// MongoDB connection string
	CONNECTION_STRING := "mongodb://localhost:27017/"

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

	// THIS CODE WORKS FOR SOME REASON...
	db.CreateCollection(context.TODO(), "test_go_collection")
	cur_collection := db.Collection("test_go_collection")
	res, err := cur_collection.InsertOne(context.TODO(), bson.D{{"name", "Alice"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted item: %v\n", res)

	// BUT THIS DOES NOT?!

	// db.CreateCollection(context.TODO(), "test_go_collection")
	// cur_collection := db.Collection("test_go_collection")
	// item1 := Game{id: 1, element: "First Go Element"}
	// item2 := Game{id: 2, element: "Second Go Element"}
	// docs := []interface{}{
	// 	item1,
	// 	item2,
	// }
	//result2, err := cur_collection.InsertMany(context.TODO(), docs)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Inserted item: %v\n", result2)

}
