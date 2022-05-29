package usecase

import (
	"context"
	"go-api/internal/domain"
	"go-api/middleware"
)

type databaseUsecase struct {
	repository domain.Repository
}

func NewUsecase(repository domain.Repository) databaseUsecase {
	return databaseUsecase{repository: repository}
}

func (usecase *databaseUsecase) GetUsers(ctx context.Context, logger *middleware.Logger,) []domain.User {
	return usecase.repository.GetUsers(ctx, logger)
} 