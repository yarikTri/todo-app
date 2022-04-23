package repository

// Authorization
type Authorization interface {
}

// TodoList
type TodoList interface {
}

// TodoItem
type TodoItem interface {
}

// Repository
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// NewRepository basicly inits new Repository
func NewRepository() *Repository {
	return &Repository{}
}
