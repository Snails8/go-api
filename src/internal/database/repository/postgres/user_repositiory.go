package postgres

import (
	"context"
	"go-api/internal/database/domain"
	"go-api/middleware"

	"github.com/jackc/pgx/v4/pgxpool"
)


type databaseRepository struct {
	dbpool *pgxpool.Pool
	logger *middleware.Logger
}

func (r *databaseRepository) GetUsers() []domain.User {
	query :=`
		SELECT users.id, users.name FROM users ORDER BY id 
	`

	rows, err := r.dbpool.Query(context.Background(), query)
	if err != nil {
		r.logger.Error.Println("query err:" + err)
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
			r.logger.Error.Println("Scan err:" + err)
		}

		users = append(users, user)
	}

	return users
}