package controllers

import (
  "encoding/json"
  "net/http"

  "github.com/johndoe/api/models"
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