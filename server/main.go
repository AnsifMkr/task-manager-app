package main

import (
	"fmt"
	"net/http"
	"todo-app/handlers"
	"todo-app/middleware"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Use(middleware.CORSMiddleware)

	r.HandleFunc("/api/todos", handlers.GetTodos).Methods("GET")
	r.HandleFunc("/api/todos", handlers.CreateTodo).Methods("POST")
	r.HandleFunc("/api/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/api/todos/{id}", handlers.DeleteTodo).Methods("DELETE")

	// Add this to handle all OPTIONS requests for CORS preflight
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("OPTIONS")

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
