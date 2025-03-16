package handler

import (
	"encoding/json"
	"net/http"

	"github.com/guneyin/todo-app/todo-api/dto"
	"github.com/guneyin/todo-app/todo-api/todo"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/todos", HandleGetTodo)
	mux.HandleFunc("POST /api/todos", HandleSubmitTodo)
}

func HandleGetTodo(w http.ResponseWriter, _ *http.Request) {
	todoSvc := todo.Service()
	todos := todoSvc.List()
	response := dto.NewTodos(todos)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(response.Bytes())
}

func HandleSubmitTodo(w http.ResponseWriter, r *http.Request) {
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
