package server

import (
	"log"
	"net/http"
)

func StartServer(port string) error {
	log.Println("Starting server...")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/todos", handleGetTodo)
	mux.HandleFunc("POST /api/todos", handleSubmitTodo)

	return http.ListenAndServe(":"+port, mux)
}

func handleGetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write([]byte(`{ "todos": ["buy some milk"] }`))
}

func handleSubmitTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
