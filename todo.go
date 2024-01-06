	package main

	import (
		"encoding/json"
		"fmt"
		"net/http"
		"github.com/rs/cors"
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

	// func createTodo(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")

	// 	var newTodo Todo
	// 	err := json.NewDecoder(r.Body).Decode(&newTodo)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusBadRequest)
	// 		return
	// 	}

	// 	newTodo.ID = len(todos) + 1
	// 	todos = append(todos, newTodo)

	// 	w.WriteHeader(http.StatusCreated)
	// 	json.NewEncoder(w).Encode(newTodo)
	// }
	func createTodo(w http.ResponseWriter, r *http.Request) {
		var todo Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Print received data
		fmt.Printf("Received Todo: %+v\n", todo)

		// Save the Todo to your data store (e.g., database)
		// ...

		// Send a response (optional)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todo)
	}

	func main() {
		router := mux.NewRouter()

		// Serve static files from the "static" directory
		router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

		// Define routes
		router.HandleFunc("/todos", getTodos).Methods("GET")
		router.HandleFunc("/todos", createTodo).Methods("POST")

		// Enable CORS
		handler := cors.Default().Handler(router)

		// Start the server
		port := 8000
		fmt.Printf("Server running on :%d\n", port)
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
		if err != nil {
			fmt.Println(err)
	
		}
	}