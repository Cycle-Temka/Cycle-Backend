package repository

import (
	"context"
	"learning-gin/internal/entities"
)

type User interface {
	Create(ctx context.Context, userCreate entities.UserCreate) (int, error)
	Get(ctx context.Context, userID int) (entities.User, error)
	GetPassword(ctx context.Context, login string) (int, string, error)
	UpdatePassword(ctx context.Context, userID int, newPassword string) error
	Delete(ctx context.Context, userID int) error
}
