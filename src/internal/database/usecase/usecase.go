package usecase

import "go-api/internal/database/domain"

type databaseUsecase struct {
	repository domain.Repository
}

func NewUsecase(repository domain.Repository) databaseUsecase {
	return databaseUsecase{repository: repository}
}

func (usecase *databaseUsecase) GetUsers() []domain.User {
	return usecase.repository.GetUsers()
} 