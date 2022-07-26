package api

import (
	"go-api/internal/testpackageA/repository/postgres"
	"go-api/internal/testpackageA/usecase"
	"go-api/middleware"
	"log"
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

func CreateUser(dbpool *pgxpool.Pool, logger *middleware.Logger) gin.HandlerFunc {
	return gin.HandlerFunc(func (c *gin.Context) {
		var userPostRequest UsersPostRequest
		if err := c.ShouldBindJSON(&userPostRequest); err != nil {
			c.JSON(http.StatusBadRequest, Error{
				Type:   "about:blank",
				Title:  _400_BAD_REQUEST_INVALID_JSON_BODY,
				Status: 400,
				Detail: "",
			})
			log.Printf("could not parse request body %v\n", err)
			return
		}
		
		repository := postgres.NewDatabaseRepository(dbpool)
		databaseUsecase := usecase.NewUserUsecase(repository)

		user, err := databaseUsecase.StoreUser(
			userPostRequest.Name, 
			userPostRequest.Email,
		)
		if err != nil {
			return
		}

		c.JSON(http.StatusNoContent, user)
	})
}