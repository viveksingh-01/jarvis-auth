package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/viveksingh-01/jarvis-auth/models"
	"github.com/viveksingh-01/jarvis-auth/utils"
)

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

	if err := utils.ValidateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// TODO
}
