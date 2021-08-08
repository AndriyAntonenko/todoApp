package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
	todoListTable   = "todo_list"
	usersListTable  = "users_list"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDb(cnfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cnfg.Host, cnfg.Port, cnfg.Username, cnfg.DBName, cnfg.Password, cnfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
