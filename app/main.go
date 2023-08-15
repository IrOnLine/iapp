package main

import (
	"fmt"
	"net/http"

<<<<<<< HEAD
	"routes"
=======
	"github.com/ironline/iapp/app"
>>>>>>> 3f3d5c15224734803c79c358214929e54f878572

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	// define routes
	routes.SetupRoutes(router)

	// start server
	fmt.Println("Starting server on port 8000")
	http.ListenAndServe(":8000", router)
}
