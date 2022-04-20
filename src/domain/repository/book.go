package repository

import (
	"context"

	"github.com/Snails8/go-api-sample/src/domain/model"
)

//　Book のI/F
//  依存性逆転の法則により infra 層は domain 層（本インターフェース）に依存
type BookRepository interface {
	GetAll(ctx context.Context) ([]*model.Book, error)
}
