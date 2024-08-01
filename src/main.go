package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// User struct defines the properties of a user
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// In-memory storage for users
var users = make(map[string]User)

// Function to add a user
func AddUser(user User) {
	users[user.ID] = user
}

// Function to get a user by ID
func GetUser(id string) (User, bool) {
	user, exists := users[id]
	return user, exists
}

// Function to delete a user by ID
func DeleteUser(id string) {
	delete(users, id)
}

// Function to update a user
func UpdateUser(id string, updatedUser User) {
	users[id] = updatedUser
}
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users/{id}", GetUserHandler).Methods("GET")
	r.HandleFunc("/users", AddUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUserHandler).Methods("DELETE")
	fmt.Println("8080")

	http.ListenAndServe(":8080", r)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user, exists := GetUser(id)
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	AddUser(user)
	w.WriteHeader(http.StatusCreated)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updatedUser User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	UpdateUser(id, updatedUser)
	w.WriteHeader(http.StatusNoContent)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	DeleteUser(id)
	w.WriteHeader(http.StatusNoContent)
}
