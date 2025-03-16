package todo_test

import (
	"github.com/guneyin/todo-app/todo-api/todo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ListItems(t *testing.T) {
	svc := todo.Service()

	item := "buy some milk"
	added, err := svc.Add(item)
	assert.NoError(t, err)
	assert.Equal(t, item, added)

	items := svc.List()
	assert.Len(t, items, 1)
	assert.Equal(t, item, items[0])
}

func Test_AddItem(t *testing.T) {
	svc := todo.Service()

	item := "buy some milk"
	added, err := svc.Add(item)
	assert.NoError(t, err)
	assert.Equal(t, item, added)
}

func Test_AddItem_Duplicate(t *testing.T) {
	svc := todo.Service()

	item := "buy some milk"
	added, err := svc.Add(item)
	assert.NoError(t, err)
	assert.Equal(t, item, added)

	duplicated, err := svc.Add(item)
	assert.Error(t, err)
	assert.Empty(t, duplicated)
}
