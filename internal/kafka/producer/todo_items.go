package producer

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/todoapps2021/telegrambot/internal/models"
)

type TodoItemsProducer struct {
	producer *kafka.Producer
	topic    *string
}

func NewTodoItemsProducer(producer *kafka.Producer) *TodoItemsProducer {
	topic := "todo_item"
	return &TodoItemsProducer{producer: producer, topic: &topic}
}

func (tip *TodoItemsProducer) Delete(userId, itemId int) error {
	_json := map[string]interface{}{
		"status":  "delete",
		"user_id": userId,
		"item_id": itemId,
	}

	valueJSON, err := json.Marshal(_json)
	if err != nil {
		return err
	}

	err = tip.producer.Produce(&kafka.Message{
		Key:   []byte(fmt.Sprint(time.Now().UTC())),
		Value: valueJSON,
		TopicPartition: kafka.TopicPartition{
			Topic:     tip.topic,
			Partition: 1, // partition number 1
		},
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

func (tip *TodoItemsProducer) Update(userId, itemId int, input models.UpdateItemInput) error {
	_json := map[string]interface{}{
		"status":  "update",
		"user_id": userId,
		"item_id": itemId,
		"item":    input,
	}

	valueJSON, err := json.Marshal(_json)
	if err != nil {
		return err
	}

	err = tip.producer.Produce(&kafka.Message{
		Key:   []byte(fmt.Sprint(time.Now().UTC())),
		Value: valueJSON,
		TopicPartition: kafka.TopicPartition{
			Topic:     tip.topic,
			Partition: 2, // partition number 2
		},
	}, nil)
	if err != nil {
		return err
	}

	return nil
}
