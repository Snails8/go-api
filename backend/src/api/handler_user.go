package api

import (
	"go-api/internal/testpackageA/repository/postgres"
	"go-api/internal/testpackageA/usecase"
	"go-api/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func GetUsers(dbpool *pgxpool.Pool, logger *middleware.Logger) gin.HandlerFunc {

	return gin.HandlerFunc(func (c *gin.Context) {
		repository := postgres.NewDatabaseRepository(dbpool)
		databaseUsecase := usecase.NewUserUsecase(repository)

		usersList := databaseUsecase.GetUsers(c, logger)
		users := make([]User, 0)
		for _, u := range usersList {
			user := User{
				Id: u.Id,
				Name: u.Name,
				Email: u.Email,
			}
			users = append(users, user)
		}

		c.JSON(http.StatusOK, GetUsersResponse{
			Users: users,
		})
	})
}