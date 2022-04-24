package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/yarikTri/todo-app"
)

// Authorization ...
type Authorization interface {
	CreateUser(user todo.User) (id int, err error)
}

// TodoList ...
type TodoList interface {
}

// TodoItem ...
type TodoItem interface {
}

// Repository ...
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// NewRepository basicly inits new Repository
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
