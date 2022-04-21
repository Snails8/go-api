package usecase

import (
	"context"

	"github.com/Snails8/go-api-sample/src/domain/model"
	"github.com/Snails8/go-api-sample/src/domain/repository"
)

type BookUseCase interface {
	GetAll(ctx context.Context) ([]*model.Book, error)
}

type bookUseCase struct {
	bookRepository repository.BookRepository
}

// Book データに関する UseCase を生成
func NewBookUseCase(br repository.BookRepository) BookUseCase {
	return &bookUseCase{
		bookRepository: br,
	}
}

func (bu bookUseCase) GetAll(ctx context.Context) (books []*model.Book, err error) {
	// Persistence（Repository）を呼出
	books, err = bu.bookRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return books, nil
}
