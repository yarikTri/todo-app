package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	todoItemsTable  = "todo_items"
	usersListsTable = "users_list"
	listsItemsTable = "lists_items"
)

// DBConfig includes all parametres required to launch Postgres DB
type DBConfig struct {
	Host     string
	Port     string
	Username string
	DBName   string
	Password string
	SSLMode  string
}

// NewPostgresDB ...
func NewPostgresDB(cfg DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
