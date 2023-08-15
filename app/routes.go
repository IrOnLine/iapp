package routes

import (
	"github.com/ironline/iapp/app/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {

	// define routes and link to controller functions

	// POST /users
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")

	// GET /users
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")

	// GET /users/{id}
	router.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")

}
