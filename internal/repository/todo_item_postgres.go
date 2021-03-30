package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/todoapps2021/telegrambot/internal/models"
)

type TodoItemPostgres struct {
	pool *pgxpool.Pool
}

func NewTodoItemPostgres(pool *pgxpool.Pool) *TodoItemPostgres {
	return &TodoItemPostgres{pool: pool}
}

func (t *TodoItemPostgres) GetAll(userId, listId int) ([]models.TodoItem, error) { // done
	ctx := context.Background()
	conn, err := t.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	items := make([]models.TodoItem, 0)
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id
									INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2 ORDER BY ti.id`,
		todoItemsTable, listsItemsTable, usersListsTable)

	rows, err := conn.Query(ctx, query, listId, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		item := models.TodoItem{}
		if err = rows.Scan(&item.Id, &item.Title, &item.Description, &item.Done); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (t *TodoItemPostgres) Delete(userId, itemId int) error { //Done
	ctx := context.Background()
	conn, err := t.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul
									WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)

	_, err = conn.Exec(ctx, query, userId, itemId)

	return err
}

func (t *TodoItemPostgres) Update(userId, itemId int, done bool) error { // done
	ctx := context.Background()
	conn, err := t.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	query := fmt.Sprintf(`UPDATE %s ti SET done=$1 FROM %s li, %s ul
									WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $2 AND ti.id = $3`,
		todoItemsTable, listsItemsTable, usersListsTable)

	_, err = conn.Exec(ctx, query, done, userId, itemId)

	return err
}
