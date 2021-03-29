package producer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/todoapps2021/telegrambot/internal/models"
)

type TodoList interface {
	Create(userId int, list models.TodoList) error
	Delete(userId, listId int) error
}

type TodoItem interface {
	Delete(userId, itemId int) error
	Update(userId, itemId int, input models.UpdateItemInput) error
}

type KProducer struct {
	TodoList
	TodoItem
}

func NewKProducer(producer *kafka.Producer) *KProducer {
	return &KProducer{
		TodoList: NewTodoListsProducer(producer),
		TodoItem: NewTodoItemsProducer(producer),
	}
}
