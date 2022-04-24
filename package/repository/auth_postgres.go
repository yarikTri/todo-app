package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/yarikTri/todo-app"
)

// AuthPostgres ...
type AuthPostgres struct {
	db *sqlx.DB
}

// NewAuthPostgres ...
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// CreateUser ...
func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil { // gets id from created row
		return 0, err
	}

	return id, nil
}
