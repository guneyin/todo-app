package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/guneyin/todo-app/todo-api/dto"
	"github.com/guneyin/todo-app/todo-api/server"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/todos", server.HandleGetTodo)
	mux.HandleFunc("POST /api/todos", server.HandleSubmitTodo)

	return httptest.NewServer(mux)
}

func Test_TodoHandler(t *testing.T) {
	ms := mockServer()
	defer ms.Close()

	todo := dto.Todo{Todo: "buy some milk"}
	payload, _ := json.Marshal(todo)

	add, err := http.Post(fmt.Sprintf("%s/api/todos", ms.URL), "application/json", bytes.NewBuffer(payload))
	assert.NoError(t, err)
	assert.NotNil(t, add)

	list, err := http.Get(fmt.Sprintf("%s/api/todos", ms.URL))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, list.StatusCode)

	todos := &dto.Todos{}
	err = json.NewDecoder(list.Body).Decode(todos)
	assert.NoError(t, err)
	assert.Equal(t, len(todos.Todos), 1)
	assert.Equal(t, todos.Todos[0], todo.Todo)
}
