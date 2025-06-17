package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/viveksingh-01/jarvis-auth/models"
	"github.com/viveksingh-01/jarvis-auth/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func SetUserCollection(c *mongo.Collection) {
	userCollection = c
}

func Register(w http.ResponseWriter, r *http.Request) {

	//  Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	// Validate the content-type
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	// Create a User struct to hold the registration data
	// This struct should match the expected JSON structure in the request body
	var user models.User

	// Decode the JSON request body into the User struct and handle any decoding errors
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the User struct fields
	if err := utils.ValidateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the username already exists in the collection
	var existingUser models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&existingUser)
	if err == nil {
		http.Error(w, "Username already exists, please try again", http.StatusBadRequest)
		return
	}
	if err != mongo.ErrNoDocuments {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// TODO
}
