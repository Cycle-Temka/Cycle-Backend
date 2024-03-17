package repository

import (
	"context"
	"learning-gin/internal/entitites"
)

type User interface {
	Create(ctx context.Context, userCreate entitites.UserCreate) (int, error)
	Get(ctx context.Context, userID int) (entitites.User, error)
	GetPassword(ctx context.Context, login string) (int, string, error)
	Delete(ctx context.Context, userID int) error
}
