package storage

import (
	"context"

	"github.com/DrSKPunk/JobBoard/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserStorage struct {
	db *pgxpool.Pool
}

func NewUserStorage(db *pgxpool.Pool) *UserStorage {
	return &UserStorage{db: db}
}

func (r *UserStorage) CreateUser(ctx context.Context, login, passwordHash string) (int, error) {
	var id int
	err := r.db.QueryRow(ctx,
		"INSERT INTO users (login, password_hash) VALUES ($1, $2) RETURNING id",
		login, passwordHash).Scan(&id)
	return id, err
}

func (r *UserStorage) GetUserByLogin(ctx context.Context, login string) (models.User, error) {
	var u models.User
	err := r.db.QueryRow(ctx,
		"SELECT id, login, password_hash FROM users WHERE login = $1",
		login).Scan(&u.ID, &u.Login, &u.Password)
	return u, err
}

func (r *UserStorage) GetUserByID(ctx context.Context, id int) (models.User, error) {
	var u models.User
	err := r.db.QueryRow(ctx,
		"SELECT id, login FROM users WHERE id = $1",
		id).Scan(&u.ID, &u.Login)
	return u, err
}
