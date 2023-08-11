package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ironline/iapp/app/models"
)

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	// get request body
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	// validate request
	if user.Email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	// create user in db
	err := models.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return 201 created
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "User created"}`))
}

// CreateGood handler
func CreateGood(w http.ResponseWriter, r *http.Request) {

	var good Good
	// decode request into good model

	err := models.CreateGood(good)
	// save good to db

}

// CreateOrder handler
func CreateOrder(w http.ResponseWriter, r *http.Request) {

	var order Order
	// decode request into order model

	err := models.CreateOrder(order)
	// save order to db

}

// GetUsers gets all users
func GetUsers(w http.ResponseWriter, r *http.Request) {

	// get users from db
	users, err := models.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return users as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUser gets a single user by id
func GetUser(w http.ResponseWriter, r *http.Request) {

	// get id from request
	vars := mux.Vars(r)
	id := vars["id"]

	// get user from db
	user, err := models.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return user as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// GetGood handler
func GetGood(w http.ResponseWriter, r *http.Request) {

	// get goodId parameter
	// fetch good from db
	// encode good as JSON response

}

// GetOrder handler
func GetOrder(w http.ResponseWriter, r *http.Request) {

	// get orderId parameter
	// fetch order from db
	// encode order as JSON response

}
