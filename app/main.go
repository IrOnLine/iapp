package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ironline/iapp/app/routes"
)

func main() {

	router := mux.NewRouter()

	// define routes
	routes.SetupRoutes(router)

	// start server
	fmt.Println("Starting server on port 8000")
	http.ListenAndServe(":8000", router)
}
