package usecase

import (
	"context"
	"fmt"
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

func (usecase *userUsecase) StoreUser(name string, email string) (domain.User, error) {
	user := domain.User{
		Name: name,
		Email: email,
	}
	
	err := usecase.repository.SaveUser(user)
	if err != nil {
		fmt.Printf("failed to create user on db: %v", err)
		return domain.User{}, err
	}

	return domain.User{}, nil
}