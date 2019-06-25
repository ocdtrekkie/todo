package main

import (
	"time"
)

const (
	// maxTitleLength is the maximum valid length of a todo item's title.
	// Todo items that exceed this length are stripped. This is to prevent
	// abuse primarily.
	maxTitleLength = 100
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
	if len(title) > maxTitleLength {
		title = title[:maxTitleLength]
	}

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
