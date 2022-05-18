package usecase

import "go-api/internal/database/domain"

type databaseUsecase struct {
	repository domain.Repository
}

func (usecase *databaseUsecase) GetUsers() []domain.User {
	return usecase.repository.GetUsers()
} 