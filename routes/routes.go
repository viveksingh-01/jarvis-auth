package routes

import (
	"github.com/gorilla/mux"
	"github.com/viveksingh-01/jarvis-auth/controllers"
)

func RegisterAuthRoutes(router *mux.Router) {
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/register", controllers.Register).Methods("POST")
}
