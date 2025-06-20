package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

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

	// Validate request method and content-type
	if err := validateRequest(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	// Generate hashed-password and store as password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword
	user.CreatedAt = time.Now()

	// Insert record in the database
	if _, err = userCollection.InsertOne(context.TODO(), user); err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}
	log.Printf("New user registered: %s", user.Username)

	// Write the response back to the client
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Registration successful. You can now log in.",
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	if err := validateRequest(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var input models.User

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		http.Error(w, "Invalid input, please check and try again.", http.StatusBadRequest)
		return
	}

	var user models.User

	// Get user from DB based on username
	err := userCollection.FindOne(context.TODO(), bson.M{"username": input.Username}).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Validate input password using hashed-password
	if !utils.ValidatePassword(input.Password, user.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT
	token, err := utils.GenerateJWT(user.ID.Hex())
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Write the response back to the client
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func validateRequest(r *http.Request) error {
	//  Check if the request method is POST
	if r.Method != http.MethodPost {
		return http.ErrNotSupported
	}
	// Validate the content-type
	if r.Header.Get("Content-Type") != "application/json" {
		return http.ErrNotSupported
	}
	return nil
}
