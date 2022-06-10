package api

import (
	"go-api/internal/testpackageA/repository/postgres"
	"go-api/internal/testpackageA/usecase"
	"go-api/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func GetTasks(dbpool *pgxpool.Pool, logger *middleware.Logger) gin.HandlerFunc {
	return gin.HandlerFunc(func (c *gin.Context)  {
		repository := postgres.NewTaskRepository(dbpool)
		usecase := usecase.NewTaskUsecase(repository)

		taskList := usecase.GetTasks(c, logger)
		tasks := make([]Task, 0)
		for _, t := range taskList {
			task := Task{
				Id: int32(t.ID),
				Title: t.Title,
				Content: t.Content,
			}
			tasks = append(tasks, task)
		}

		c.JSON(http.StatusOK, GetTasksResponse{
			Tasks: tasks,
		})
	})
}