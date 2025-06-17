package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/viveksingh-01/jarvis-auth/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectToDB() {
	// Create a new context with a timeout for the MongoDB connection
	// This context will be used to set a timeout for the connection attempt
	// If the connection takes longer than the timeout, it will be cancelled
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Ensure that the context is cancelled to free up resources
	// after the connection attempt is complete
	defer cancel()

	// Connect to the MongoDB database using the URI from environment variables
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}

	// Log a message indicating successful connection
	log.Println("Connected to MongoDB successfully")

	// Set the global DB variable to the connected client
	// This allows other parts of the application to access the database
	// The client.Database method is used to specify the database name
	DB = client.Database("jarvisdb")

	if DB == nil {
		log.Fatal("Database connection is not initialized")
	}
	controllers.SetUserCollection(DB.Collection("users"))
}
