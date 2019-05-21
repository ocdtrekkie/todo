package main

import (
	"time"
)

// Todo ...
type Todo struct {
	ID        uint64
	Done      bool
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func newTodo(title string) *Todo {
	return &Todo{
		Title:     title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (t *Todo) setTitle(title string) {
	t.Title = title
	t.UpdatedAt = time.Now()
}

func (t *Todo) toggleDone() {
	t.Done = !t.Done
	t.UpdatedAt = time.Now()
}

// TodoList ...
type TodoList []*Todo

func (a TodoList) Len() int           { return len(a) }
func (a TodoList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a TodoList) Less(i, j int) bool { return a[i].ID < a[j].ID }
