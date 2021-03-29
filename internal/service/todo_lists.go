package service

import (
	"github.com/todoapps2021/telegrambot/internal/kafka/producer"
	"github.com/todoapps2021/telegrambot/internal/models"
	"github.com/todoapps2021/telegrambot/internal/repository"
)

type TodoListService struct {
	repo     repository.TodoList
	producer producer.TodoList // TODO
}

func NewTodoListService(repo repository.TodoList, producer producer.TodoList) *TodoListService {
	return &TodoListService{repo: repo, producer: producer}
}

func (s *TodoListService) Create(userId int, list models.TodoList) error {

	// return s.repo.Create(userId, list)     // to postgres sql
	return s.producer.Create(userId, list) // to kafka
}

func (s *TodoListService) GetAll(userId int) ([]models.TodoList, error) { // Done
	return s.repo.GetAll(userId)
}

func (s *TodoListService) Delete(userId, listId int) error {
	// return s.repo.Delete(userId, listId) // to postgres sql
	return s.producer.Delete(userId, listId) // to kafka
}
