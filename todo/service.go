package todo

import (
	"errors"
)

type Service interface {
	GetTodos(UserID int) ([]Todo, error)
	GetTodoById(input GetTodoDetailInput, UserID int) (Todo, error)
	CreateTodo(input CreateTodoInput) (Todo, error)
	UpdateTodo(inputID GetTodoDetailInput, inputData UpdateTodoInput) (Todo, error)
	DeleteTodo(inputID GetTodoDetailInput, UserID int) (Todo, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTodos(UserID int) ([]Todo, error) {
	//! melakukan switching percabangan
	if UserID != 0 {
		todos, err := s.repository.FindByUserID(UserID)
		if err != nil {
			return todos, err
		}
		return todos, nil
	} else {
		todos, err := s.repository.FindAll()
		if err != nil {
			return todos, err
		}
		return todos, nil
	}
}

func (s *service) GetTodoById(input GetTodoDetailInput, UserID int) (Todo, error) {
	if UserID != 0 {
		todo, err := s.repository.FindByUserIDAndTodoId(input.ID, UserID)
		if err != nil {
			return todo, err
		}
		return todo, nil
	} else {
		todo, err := s.repository.FindByID(input.ID)
	
		if err != nil {
			return todo, err
		}
	
		return todo, nil
	}
}

func (s *service) CreateTodo(input CreateTodoInput) (Todo, error) {
	todo := Todo{}
	todo.Task = input.Task
	todo.Title = input.Title
	todo.IsCompleted = 0
	todo.UserID = input.User.ID

	newTodo, err := s.repository.SaveTodo(todo)
	if err != nil {
		return newTodo, err
	}

	return newTodo, nil
}

func(s *service) UpdateTodo(inputID GetTodoDetailInput, inputData UpdateTodoInput) (Todo, error) {
	todo, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return todo, err
	}

	if todo.UserID != inputData.User.ID {
		return todo, errors.New("not owner of todo")
	}

	todo.Task = inputData.Task
	todo.Title = inputData.Title
	todo.IsCompleted = inputData.IsCompleted

	updated, err := s.repository.UpdateTodo(todo)
	if err != nil {
		return updated, err
	}

	return updated, nil
}

func(s *service) DeleteTodo(inputID GetTodoDetailInput, UserID int) (Todo, error) {
	todo, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return todo, err
	}

	if todo.UserID != UserID {
		return todo, errors.New("not owner of todo")
	}

	deleted, err := s.repository.DeleteTodo(todo)
	if err != nil {
		return deleted, err
	}

	return deleted, nil
}