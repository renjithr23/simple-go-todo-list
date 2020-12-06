package main

import (
	"errors"
)

type StorageMemory struct {
	todos []Todo
}

func (s *StorageMemory) GetTodo(id int) (Todo, error) {

	for _, todo := range s.todos {
		if todo.Id == id {
			return todo, nil
		}
	}

	return Todo{}, errors.New("Todo Not found")

}

func (s *StorageMemory) DeleteTodo(id int) error {

	var exists bool = false

	for j, todo := range s.todos {
		if todo.Id == id {
			exists = true
			copy(s.todos[j:], s.todos[j+1:])
			s.todos = s.todos[:len(s.todos)-1]
			break
		}
	}

	if !exists {
		return errors.New("Todo Not found")
	}

	return nil

}

func (s *StorageMemory) CreateTodo(todo Todo) error {

	s.todos = append(s.todos, todo)

	return nil

}

func (s *StorageMemory) UpdateTodo(update_todo Todo) error {

	var exists bool = false

	for j, todo := range s.todos {
		if todo.Id == update_todo.Id {
			exists = true
			s.todos[j] = update_todo
			break
		}
	}

	if !exists {
		return errors.New("Todo Not found")
	}

	return nil

}

func (s *StorageMemory) ListTodo() ([]Todo, error) {

	return s.todos, nil

}
