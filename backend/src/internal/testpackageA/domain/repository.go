package domain

import (
	"context"
	"go-api/middleware"
)

type Repository interface {
	GetUsers(ctx context.Context, logger *middleware.Logger) []User
}

type TaskRepository interface {
	GetTasks(ctx context.Context, logger *middleware.Logger) []Task
}