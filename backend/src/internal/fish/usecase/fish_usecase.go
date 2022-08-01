package usecase

import (
	"context"
	"go-api/internal/fish/domain"
	"go-api/middleware"
)

type fishUsecase struct {
	repository domain.FishRepository
}

func NewFishUsecase(repository domain.FishRepository) fishUsecase {
	return fishUsecase{repository: repository}
}

func (u *fishUsecase) GetFishes(ctx context.Context, logger *middleware.Logger) []domain.Fish {
	return u.repository.GetFishes(ctx, logger)
}