package usecase

import (
	"context"
	"go-api/internal/testpackageA/domain"
	"go-api/middleware"
)

type databaseUsecase struct {
	repository domain.Repository
}

type taskUsecase struct {
	repository domain.TaskRepository
}

func NewUsecase(repository domain.Repository) databaseUsecase {
	return databaseUsecase{repository: repository}
}

func NewTaskUsecase(repository domain.TaskRepository) taskUsecase {
	return taskUsecase{repository: repository}
}

func (usecase *databaseUsecase) GetUsers(ctx context.Context, logger *middleware.Logger,) []domain.User {
	return usecase.repository.GetUsers(ctx, logger)
} 

func (usecase *taskUsecase) GetTasks(ctx context.Context, logger *middleware.Logger) []domain.Task {
	return usecase.repository.GetTasks(ctx, logger)
}