package postgres

import (
	"context"
	"go-api/internal/domain"
	"go-api/middleware"

	"github.com/jackc/pgx/v4/pgxpool"
)

type taskRepository struct {
	dbpool *pgxpool.Pool
}

func NewTaskRepository(dbpool *pgxpool.Pool) domain.TaskRepository {
	return &taskRepository{
		dbpool: dbpool,
	}
}

func (r *taskRepository) GetTasks(ctx context.Context, logger *middleware.Logger) []domain.Task {
	query := 
	`
		SELECT
			title
			, content
			, created_at
			, updated_at
		FROM
			tasks
		ORDER BY id	
	`

	rows, err := r.dbpool.Query(context.Background(), query)
	if err != nil {
		logger.Error.Printf("could not exec get tasks query: %\n", err)
		return []domain.Task{}
	}

	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		task := domain.Task{}

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Content,
		)
		if err != nil {
			logger.Error.Printf("could not exec get tasks query: %\n", err)
			return []domain.Task{}
		}

		tasks = append(tasks, task)
	}
	return tasks
}