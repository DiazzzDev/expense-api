package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"example/expense-api/internal/apperr"
	"example/expense-api/internal/model"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Ensure(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO users (id) VALUES ($1) ON CONFLICT (id) DO NOTHING`, id)
	if err != nil {
		return fmt.Errorf("ensure user: %w", err)
	}
	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*model.User, error) {
	var u model.User
	err := r.db.QueryRow(ctx,
		`SELECT id, display_name, created_at, updated_at FROM users WHERE id = $1`, id,
	).Scan(&u.ID, &u.DisplayName, &u.CreatedAt, &u.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, apperr.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}
	return &u, nil
}

func (r *UserRepository) UpdateDisplayName(ctx context.Context, id, name string) (*model.User, error) {
	var u model.User
	err := r.db.QueryRow(ctx,
		`UPDATE users SET display_name = $2, updated_at = now()
		 WHERE id = $1
		 RETURNING id, display_name, created_at, updated_at`, id, name,
	).Scan(&u.ID, &u.DisplayName, &u.CreatedAt, &u.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, apperr.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("update user: %w", err)
	}
	return &u, nil
}
