package postgres

import (
	"context"
	"go-api/internal/database/domain"
	"go-api/middleware"

	"github.com/jackc/pgx/v4/pgxpool"
)


type databaseRepository struct {
	dbpool *pgxpool.Pool
}

// handler などNewする際に使用
func NewDatabaseRepository(dbpool *pgxpool.Pool,) domain.Repository {
	return &databaseRepository{
		dbpool: dbpool,
	}
}

func (r *databaseRepository) GetUsers(ctx context.Context, logger *middleware.Logger) []domain.User {
	query :=`
		SELECT users.id, users.name FROM users ORDER BY id 
	`

	rows, err := r.dbpool.Query(context.Background(), query)
	if err != nil {
		logger.Error.Printf("could not query: %v\n", err)
		return []domain.User{}
	}

	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}

		err := rows.Scan(
			&user.ID,
			&user.Name,
		)
		if err != nil {
			logger.Error.Printf("could not query: %v\n", err)
			return []domain.User{}
		}

		users = append(users, user)
	}

	return users
}