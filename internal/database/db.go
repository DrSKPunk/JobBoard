package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() error {
	dsn := "postgres://jobboard_user:jobboard_password@127.0.0.1:5433/jobboard_db?sslmode=disable"

	var err error
	DB, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		return fmt.Errorf("не удалось создать пул подключений: %w", err)
	}

	if err := DB.Ping(context.Background()); err != nil {
		return fmt.Errorf("база данных не отвечает: %w", err)
	}

	fmt.Println("Успешное подключение к PostgreSQL!")
	return nil
}
