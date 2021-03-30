package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/todoapps2021/telegrambot/internal/models"
)

type TodoListPostgres struct {
	pool *pgxpool.Pool
}

func NewTodoListPostgres(pool *pgxpool.Pool) *TodoListPostgres {
	return &TodoListPostgres{pool: pool}
}

func (t *TodoListPostgres) Create(userId int, list models.TodoList) error { //Done
	ctx := context.Background()
	conn, err := t.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(ctx, createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		if e := tx.Rollback(ctx); e != nil {
			return err
		}
		return err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(ctx, createUsersListQuery, userId, id)
	if err != nil {
		if e := tx.Rollback(ctx); e != nil {
			return err
		}
		return err
	}

	return tx.Commit(ctx)
}

func (t *TodoListPostgres) GetAll(userId int) ([]models.TodoList, error) { //Done
	ctx := context.Background()
	conn, err := t.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	lists := make([]models.TodoList, 0)
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 ORDER BY tl.id",
		todoListsTable, usersListsTable)

	rows, err := conn.Query(ctx, query, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		list := models.TodoList{}
		if err = rows.Scan(&list.Id, &list.Title, &list.Description); err != nil {
			return nil, err
		}

		lists = append(lists, list)
	}

	return lists, nil
}

func (t *TodoListPostgres) GetById(telegramId, listId int) (models.TodoList, error) {
	ctx := context.Background()
	var list models.TodoList
	conn, err := t.pool.Acquire(ctx)
	if err != nil {
		return list, err
	}
	defer conn.Release()

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl
								INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = (SELECT id from users WHERE telegram_id = %v) AND ul.list_id = $1`,
		todoListsTable, usersListsTable, telegramId)

	row := conn.QueryRow(ctx, query, listId)
	if err := row.Scan(&list.Id, &list.Title, &list.Description); err != nil {
		return list, err
	}

	return list, nil
}

func (t *TodoListPostgres) Delete(userId, listId int) error { // done
	ctx := context.Background()
	conn, err := t.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id = $2",
		todoListsTable, usersListsTable)

	_, err = conn.Exec(ctx, query, userId, listId)

	return err
}
