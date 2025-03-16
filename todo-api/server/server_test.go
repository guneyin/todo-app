package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/guneyin/todo-app/todo-api/dto"
	"github.com/guneyin/todo-app/todo-api/handler"
	"github.com/stretchr/testify/assert"
)

func mockServer() *httptest.Server {
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	return httptest.NewServer(mux)
}

func Test_TodoHandler(t *testing.T) {
	ms := mockServer()
	defer ms.Close()

	todo := dto.Todo{Todo: "buy some milk"}
	payload, _ := json.Marshal(todo)

	add, err := http.Post(fmt.Sprintf("%s/api/todos", ms.URL), "application/json", bytes.NewBuffer(payload))
	require.NoError(t, err)
	assert.NotNil(t, add)

	list, err := http.Get(fmt.Sprintf("%s/api/todos", ms.URL))
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, list.StatusCode)

	todos := &dto.Todos{}
	err = json.NewDecoder(list.Body).Decode(todos)
	require.NoError(t, err)
	assert.Len(t, todos.Todos, 1)
	assert.Equal(t, todos.Todos[0], todo.Todo)
}
