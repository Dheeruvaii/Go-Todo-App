package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var todos []Todo

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newTodo Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newTodo.ID = len(todos) + 1
	todos = append(todos, newTodo)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

func main() {
	router := mux.NewRouter()
	todos = append(todos, Todo{ID: 1, Title: "todo 1"})

	// Define routes
	router.HandleFunc("/todos", getTodos).Methods("GET")
	router.HandleFunc("/todos", createTodo).Methods("POST")
	
    // Enable CORS
	// Start the server
	port := 8000
	fmt.Printf("Server running on :%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	if err != nil {
		fmt.Println(err)
	}
}
