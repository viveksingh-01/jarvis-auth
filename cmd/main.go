package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Welcome to JARVIS authentication system.")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading the .env file.")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server started at port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
