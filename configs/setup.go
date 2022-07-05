package configs

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017/gingonic" // DB connection string

func ConnectDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI(connectionString) // Set client options
	client, err := mongo.Connect(context.TODO(), clientOptions)  // connect with MongoDB
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil) // Check the connection
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

//Client instance
var DB *mongo.Client = ConnectDB()

//getting database collections instance
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("gingonic").Collection(collectionName)
	return collection
}
