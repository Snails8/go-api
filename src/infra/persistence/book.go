package persistence

// このパッケージ名いけてない

import (
	"context"
	"github.com/Snails8/go-api-sample/src/domain/repository"
	"time"
)

type bookPersistence struct{}

//  Book データに関する Persistence を生成
func NewBookPersistence() repository.BookRepository {
	return &bookPersistence{}
}

// データを取得（簡略化のためにモックデータを使用)
func (bp bookPersistence) GetAll(ctx context.Context) ([]*model.Book, error) {
	book1 := model.book{}
	book1.Id = 1
	book1.Title = "DDDの本"
	book1.Author = "クレクレくん"
	book1.IssuedAt = time.Now().Add(-24 * time.Hour)

	book2 := model.Book{}
	book2.Id = 2
	book2.Title = "レイヤードアーキテクチャが分かる本"
	book2.Author = "はなこさん"
	book2.IssuedAt = time.Now().Add(-24 * 7 * time.Hour)

	return []*model.Book{&book1, &book2}, nil
}
