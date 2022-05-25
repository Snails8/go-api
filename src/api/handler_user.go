package api

import (
	"go-api/internal/database/repository/postgres"
	"go-api/internal/database/usecase"
	"go-api/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func GetUsers(dbpool *pgxpool.Pool, logger *middleware.Logger) gin.HandlerFunc {

	return gin.HandlerFunc(func (c *gin.Context) {
		repository := postgres.NewDatabaseRepository(dbpool)
		databaseUsecase := usecase.NewUsecase(repository)

		usersList := databaseUsecase.GetUsers(c, logger)
		users := make([]User, 0)
		for _, u := range usersList {
			user := User{
				Id: u.ID,
				Name: u.Name,
			}
			users = append(users, user)
		}

		c.JSON(http.StatusOK, GetUsersResponse{
			Users: users,
		})
	})
}