package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/todoapps2021/telegrambot/internal/models"
)

type Authorization interface {
	Login(username, password string, telegramId int) error
	CheckUser(username, password string) (bool, error)
	GetUserIdByTgId(telegram_id int) (int, error)
}

type TodoList interface {
	Create(userId int, list models.TodoList) error
	GetAll(userId int) ([]models.TodoList, error)
	GetById(telegramId, listId int) (models.TodoList, error)
	Delete(userId, listId int) error
}

type TodoItem interface {
	GetAll(userId, listId int) ([]models.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, done bool) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(pool),
		TodoList:      NewTodoListPostgres(pool),
		TodoItem:      NewTodoItemPostgres(pool),
	}
}
