package postgres

import (
	"context"
	"go-api/internal/fish/domain"
	"go-api/middleware"

	"github.com/jackc/pgx/v4/pgxpool"
)

type fishRepository struct {
	dbpool *pgxpool.Pool
}

func NewFishRepository(dbpool *pgxpool.Pool) domain.FishRepository {
	return &fishRepository{dbpool: dbpool}
}

func (r *fishRepository) GetFishes(ctx context.Context, logger *middleware.Logger) []domain.Fish {
	query := `
		SELECT fishes.id, fishes.name FROM fishes ORDER BY id 
	`

	rows, err := r.dbpool.Query(context.Background(), query)
	if err != nil {
		logger.Error.Panicf("could not query: %v\n", err)
	}
	defer rows.Close()

	var fishes []domain.Fish
	for rows.Next() {
		fish := domain.Fish{}
		err := rows.Scan(
			&fish.Id,
			&fish.Name,
		)
		if err != nil {
			logger.Error.Printf("could not query: %v\n", err)
			return []domain.Fish{}
		}
		fishes = append(fishes, fish)
	}

	return fishes
}