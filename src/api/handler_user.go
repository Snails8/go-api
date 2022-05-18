package api

import (
	"go-api/internal/database/repository/postgres"
	"go-api/internal/database/usecase"
	"go-api/middleware"

	"github.com/gin-gonic/gin"
	"google.golang.org/genproto/googleapis/spanner/admin/database/v1"
)

func GetUsers(dbpool *dbpool.Pool, logger *middleware.Logger) gin.HandlerFunc {

	return gin.HandlerFunc(func (c *gin.Context) {
		repository := postgres.NewDatabaseRepository(dbpool, logger)
		databaseUsecase := usecase.NewUsecase(repository)

		usersList := databaseUsecase.GetUsers()
		users := make([]User, 0)
		for _, u := range usersList {
			user := User{
				Id: u.ID,
				Name: u.Name,
			}
			users = append(users, user)
		}

		c.Json
	})
}