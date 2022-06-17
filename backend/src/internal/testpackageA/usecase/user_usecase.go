package usecase

import (
	"context"
	"go-api/internal/testpackageA/domain"
	"go-api/middleware"
)

type userUsecase struct {
	repository domain.Repository
}

func NewUserUsecase(repository domain.Repository) userUsecase {
	return userUsecase{repository: repository}
}

func (usecase *userUsecase) GetUsers(ctx context.Context, logger *middleware.Logger,) []domain.User {
	return usecase.repository.GetUsers(ctx, logger)
} 

