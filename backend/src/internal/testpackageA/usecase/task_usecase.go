
package usecase

import (
	"context"
	"go-api/internal/testpackageA/domain"
	"go-api/middleware"
)

type taskUsecase struct {
	repository domain.TaskRepository
}


func NewTaskUsecase(repository domain.TaskRepository) taskUsecase {
	return taskUsecase{repository: repository}
}

func (usecase *taskUsecase) GetTasks(ctx context.Context, logger *middleware.Logger) []domain.Task {
	return usecase.repository.GetTasks(ctx, logger)
}