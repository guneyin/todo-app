package dto

import "encoding/json"

type Todos struct {
	Todos []string `json:"todos"`
}

type Todo struct {
	Todo string `json:"todo"`
}

func NewTodos(sl []string) *Todos {
	return &Todos{Todos: sl}
}

func (t *Todos) Bytes() []byte {
	b, _ := json.Marshal(t)
	return b
}
