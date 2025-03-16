package server

import (
	"encoding/json"
	"github.com/guneyin/todo-app/todo-api/dto"
	"github.com/guneyin/todo-app/todo-api/todo"
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
	todoSvc := todo.Service()
	todos := todoSvc.List()
	response := dto.NewTodos(todos)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(response.Bytes())
}

func handleSubmitTodo(w http.ResponseWriter, r *http.Request) {
	todoSvc := todo.Service()

	todoItem := &dto.Todo{}
	err := json.NewDecoder(r.Body).Decode(todoItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = todoSvc.Add(todoItem.Todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
