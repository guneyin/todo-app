package todo_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/guneyin/todo-app/todo-api/todo"
	"github.com/stretchr/testify/assert"
)

func Test_ListItems(t *testing.T) {
	t.Run("list todo items", func(t *testing.T) {
		svc := todo.NewService()

		item := "buy some milk"
		added, err := svc.Add(item)
		require.NoError(t, err)
		assert.Equal(t, item, added)

		items := svc.List()
		assert.Len(t, items, 1)
		assert.Equal(t, item, items[0])
	})

}

func Test_AddItem(t *testing.T) {
	t.Run("add item", func(t *testing.T) {
		svc := todo.NewService()

		item := "buy some milk"
		added, err := svc.Add(item)
		require.NoError(t, err)
		assert.Equal(t, item, added)
	})
}

func Test_AddItem_Duplicate(t *testing.T) {
	t.Run("add duplicate item", func(t *testing.T) {
		svc := todo.NewService()

		item := "buy some milk"
		added, err := svc.Add(item)
		require.NoError(t, err)
		assert.Equal(t, item, added)

		duplicated, err := svc.Add(item)
		require.Error(t, err)
		assert.Empty(t, duplicated)
	})
}
