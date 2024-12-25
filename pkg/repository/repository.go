package repository

import (
	"github.com/UmalatDukuev/todo"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	// GenerateToken(username, password string) (string, error)
	// ParseToken(token string) (int, error)
}

type TodoList interface {
	// Create(userId int, list todo.TodoList) (int, error)
	// GetAll(userId int) ([]todo.TodoList, error)
	// GetById(userId, listId int) (todo.TodoList, error)
	// Delete(userId, listId int) error
	// Update(userId, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
	// Create(userId, listId int, item todo.TodoItem) (int, error)
	// GetAll(userId, listId int) ([]todo.TodoItem, error)
	// GetById(userId, itemId int) (todo.TodoItem, error)
	// Delete(userId, itemId int) error
	// Update(userId, itemId int, input todo.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}

// func NewRepository(db *sqlx.DB) *Repository {
// 	return &Repository{
// 		Authorization: NewAuthPostgres(db),
// 		TodoList:      NewTodoListPostgres(db),
// 		TodoItem:      NewTodoItemPostgres(db),
// 	}
// }
