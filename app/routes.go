package routes

import (
<<<<<<< HEAD
	"controllers"
=======
	"github.com/ironline/iapp/app"
>>>>>>> 3f3d5c15224734803c79c358214929e54f878572

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
