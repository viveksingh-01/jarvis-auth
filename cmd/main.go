package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/viveksingh-01/jarvis-auth/config"
	"github.com/viveksingh-01/jarvis-auth/routes"
)

func main() {
	fmt.Println("Welcome to JARVIS authentication system.")

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading the .env file.")
	}

	// Establish connection to DB
	config.ConnectToDB()

	// Initialize the router
	r := mux.NewRouter()
	routes.RegisterAuthRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server started at port", port)
	// Start the server with the router and port
	log.Fatal(http.ListenAndServe(port, r))
}
