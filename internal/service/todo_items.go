package service

import (
	"github.com/todoapps2021/telegrambot/internal/kafka/producer"
	"github.com/todoapps2021/telegrambot/internal/models"
	"github.com/todoapps2021/telegrambot/internal/repository"
)

type TodoItemService struct {
	itemRepo repository.TodoItem
	producer producer.TodoItem // TODO
}

func NewTodoItemService(itemRepo repository.TodoItem, producer producer.TodoItem) *TodoItemService {
	return &TodoItemService{itemRepo: itemRepo, producer: producer}
}

func (s *TodoItemService) GetAll(userId, listId int) ([]models.TodoItem, error) {
	return s.itemRepo.GetAll(userId, listId)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	// return s.itemRepo.Delete(userId, itemId) // to postgres sql
	return s.producer.Delete(userId, itemId) // to kafka
}

func (s *TodoItemService) Update(userId, itemId int, done bool) error {
	var input models.UpdateItemInput
	input.Done = &done
	// return s.itemRepo.Update(userId, itemId, done)
	return s.producer.Update(userId, itemId, input)
}
