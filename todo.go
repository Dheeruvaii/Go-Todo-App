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

func main() {
	router := mux.NewRouter()
	todos = append(todos, Todo{ID: 1, Title: "todo 1"})

	// Define route
	router.HandleFunc("/todos", getTodos).Methods("GET")

	// Start the server
	port := 8000
	fmt.Printf("Server running on :%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	if err != nil {
		fmt.Println(err)
	}
}
