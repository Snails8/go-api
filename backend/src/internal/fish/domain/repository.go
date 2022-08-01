package domain

import (
	"context"
	"go-api/middleware"
)

type FishRepository interface {
	GetFishes(ctx context.Context, logger *middleware.Logger) []Fish
}