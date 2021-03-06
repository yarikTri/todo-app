package service

import (
	"github.com/yarikTri/todo-app"
	"github.com/yarikTri/todo-app/package/repository"
)

// Authorization declares what Auth supposed to implement
type Authorization interface {
	CreateUser(user todo.User) (id int, err error)
}

// TodoList declares what List supposed to implement
type TodoList interface {
}

// TodoItem declares what Item supposed to implement
type TodoItem interface {
}

// Service connects all services together
type Service struct {
	Authorization
	TodoList
	TodoItem
}

// NewService basicly inits new Service
func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
