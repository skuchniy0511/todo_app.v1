package repository

import (
	app "app"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user app.User) (int, error)
	GetUser(username, password string) (app.User, error)
}

type TodoList interface {
	Create(userId int, list app.TodoList) (int, error)
	GetAll(UserId int) ([]app.TodoList, error)
	GetById(userId, listId int) (app.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input app.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item app.TodoItem) (int, error)
	GetAll(UserId, listId int) ([]app.TodoItem, error)
	GetById(userId, itemId int) (app.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input app.UpdateItemsInput) error
}
type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
