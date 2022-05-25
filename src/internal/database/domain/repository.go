package domain

import (
	"context"
	"go-api/middleware"
)

type Repository interface {
	GetUsers(ctx context.Context, logger *middleware.Logger) []User
}