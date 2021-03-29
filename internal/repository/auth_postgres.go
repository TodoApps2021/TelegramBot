package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthPostgres struct {
	pool *pgxpool.Pool
}

func NewAuthPostgres(pool *pgxpool.Pool) *AuthPostgres {
	return &AuthPostgres{pool: pool}
}

func (as *AuthPostgres) Login(username, password string, telegramId int) error {
	ctx := context.Background()
	query := fmt.Sprintf("UPDATE %s SET telegram_id=%v WHERE username=$1 AND password_hash=$2", usersTable, telegramId)

	conn, err := as.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	if ok, err := as.CheckUser(username, password); !ok || err != nil {
		return errors.New("This user not exists.")
	}

	_, err = conn.Exec(ctx, query, username, password)

	return err
}

func (as *AuthPostgres) CheckUser(username, password string) (bool, error) {
	ctx := context.Background()
	query := fmt.Sprintf("SELECT EXISTS(SELECT * FROM %s WHERE username=$1 AND password_hash=$2)", usersTable)
	var isExists bool

	conn, err := as.pool.Acquire(ctx)
	if err != nil {
		return false, err
	}
	defer conn.Release()

	row := conn.QueryRow(ctx, query, username, password)
	if err := row.Scan(&isExists); err != nil {
		return false, err
	}

	return isExists, nil
}

func (as *AuthPostgres) GetUserIdByTgId(telegram_id int) (int, error) {
	ctx := context.Background()
	query := fmt.Sprintf("SELECT id FROM %s WHERE telegram_id=%v", usersTable, telegram_id)

	conn, err := as.pool.Acquire(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Release()

	var userId = 0
	row := conn.QueryRow(ctx, query)
	if err := row.Scan(&userId); err != nil {
		return 0, err
	}

	return userId, nil
}
