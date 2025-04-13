package main

import (
	"errors"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CompletedAt *time.Time
}

type TodoList []Todo

func (todoList *TodoList) Add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todoList = append(*todoList, todo)
}

// Helper to find if the todo exists or not
func (todoList *TodoList) doesTodoExists(id int) error {
	todo := (*todoList)[id]

	if todo.Title == "" {
		err := errors.New("invalid todo id")
		return err
	}

	return nil
}

func (todoList *TodoList) Delete(id int) error {
	t := *todoList

	if err := t.doesTodoExists(id); err != nil {
		return err
	}

	*todoList = slices.Delete(t, id, id+1)

	return nil
}

func (todoList *TodoList) toggleCompleted(id int) error {
	t := *todoList

	if err := t.doesTodoExists(id); err != nil {
		return err
	}

	isTodoCompleted := t[id].Completed

	if !isTodoCompleted {
		completedTime := time.Now()
		t[id].CompletedAt = &completedTime
	}

	t[id].Completed = !isTodoCompleted

	return nil
}

func (todoList *TodoList) Update(id int, title string) error {
	t := *todoList

	if err := t.doesTodoExists(id); err != nil {
		return err
	}

	t[id].Title = title
	t[id].UpdatedAt = time.Now()

	return nil
}

func (todoList *TodoList) Get(id int) (Todo, error) {
	t := *todoList

	if err := t.doesTodoExists(id); err != nil {
		return Todo{}, err
	}

	return t[id], nil
}

func (todoList *TodoList) Print() {
	t := table.New(os.Stdout)

	t.SetRowLines(false)
	t.AddHeaders("ID", "Task", "Completed", "Created At", "Updated At", "Completed At")

	for i, todo := range *todoList {
		completed := "X"
		completedAt := ""

		if todo.Completed {
			completed = "O"

			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.RFC1123)
			}
		}

		t.AddRow(strconv.Itoa(i), todo.Title, completed, todo.CreatedAt.Format(time.RFC1123), todo.UpdatedAt.Format(time.RFC1123), completedAt)
	}

	t.Render()
}
