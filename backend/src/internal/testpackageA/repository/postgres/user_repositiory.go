package postgres

import (
	"context"
	"go-api/internal/testpackageA/domain"
	"go-api/middleware"
	"log"
	"time"

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
		SELECT users.id, users.name, users.email FROM users ORDER BY id 
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
			&user.Id,
			&user.Name,
			&user.Email,
		)
		if err != nil {
			logger.Error.Printf("could not query: %v\n", err)
			return []domain.User{}
		}

		users = append(users, user)
	}

	return users
}

func (r databaseRepository) SaveUser(user domain.User) error {
	query := `INSERT INTO users (name, email, created_at, updated_at) 
				VALUES ($1, $2, $3, $3)
				ON CONFLICT (email)
				DO UPDATE SET name=$1, email=$2, updated_at=$3
			`

	_, err := r.dbpool.Exec(context.Background(), query, user.Name, user.Email, time.Now())
	if err != nil {
		log.Printf("error: user exec query:  %v\n", err)
	} 
	return err
}
