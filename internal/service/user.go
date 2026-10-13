package service

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf8"

	"example/expense-api/internal/apperr"
	"example/expense-api/internal/model"
)

type UserRepository interface {
	Ensure(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*model.User, error)
	UpdateDisplayName(ctx context.Context, id, name string) (*model.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Ensure(ctx context.Context, id string) error {
	return s.repo.Ensure(ctx, id)
}

func (s *UserService) Get(ctx context.Context, id string) (*model.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) UpdateDisplayName(ctx context.Context, id, name string) (*model.User, error) {
	name = strings.TrimSpace(name)
	if n := utf8.RuneCountInString(name); n < 1 || n > 60 {
		return nil, fmt.Errorf("%w: display_name must be 1-60 characters", apperr.ErrInvalid)
	}
	return s.repo.UpdateDisplayName(ctx, id, name)
}
