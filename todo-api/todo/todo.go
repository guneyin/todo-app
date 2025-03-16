package todo

import (
	"errors"
	"maps"
	"slices"
	"sync"
)

var ErrItemAlreadyExist = errors.New("item already exists")

var once sync.Once

type Todo struct {
	items map[string]struct{}
}

var svc *Todo

func Service() *Todo {
	once.Do(func() {
		svc = &Todo{
			items: make(map[string]struct{}),
		}
	})

	return svc
}

func (t *Todo) List() []string {
	return slices.Collect(maps.Keys(t.items))
}

func (t *Todo) Add(item string) (string, error) {
	if _, ok := t.items[item]; ok {
		return "", ErrItemAlreadyExist
	}

	t.items[item] = struct{}{}
	return item, nil
}
