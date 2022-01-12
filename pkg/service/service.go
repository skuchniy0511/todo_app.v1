package service

import (
	"app/pkg/repository"

	app "app"
)

type Authorization interface {
	CreateUser(user app.User) (int, error)
	GenerateToken(Username, Password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list app.TodoList) (int, error)
	GetAll(UserId int) ([]app.TodoList, error)
	GetById(userId, listId int) (app.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input app.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item app.TodoItem) (int, error)
	GetAll(UserId, listId int) ([]app.TodoItem, error)
	GetById(userId, itemId int) (app.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input app.UpdateItemsInput) error
}
type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
