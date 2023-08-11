package main

import (
	"fmt"
	"net/http"

	"routes"

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
