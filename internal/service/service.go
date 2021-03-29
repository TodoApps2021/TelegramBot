package service

import (
	"github.com/todoapps2021/telegrambot/internal/kafka/producer"
	"github.com/todoapps2021/telegrambot/internal/models"
	"github.com/todoapps2021/telegrambot/internal/repository"
)

type Authorization interface {
	Login(username, password string, telegramId int) error
	CheckUser(username, password string) (bool, error)
	GetUserIdByTgId(telegram_id int) (int, error)
}

type TodoList interface {
	Create(userId int, list models.TodoList) error // todo
	GetAll(userId int) ([]models.TodoList, error)
	Delete(userId, listId int) error //
}

type TodoItem interface {
	GetAll(userId, listId int) ([]models.TodoItem, error) //
	Delete(userId, itemId int) error                      //
	Update(userId, itemId int, done bool) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository, producer *producer.KProducer) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList, producer.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, producer.TodoItem),
	}
}
