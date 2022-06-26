package api

import (
	// "go-api/internal/testpackageA/repository/postgres"
	// "go-api/internal/testpackageA/usecase"
	"go-api/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

type test struct {
	test string
}

func Dashboard(dbpool *pgxpool.Pool, logger *middleware.Logger) gin.HandlerFunc {

	return gin.HandlerFunc(func (c *gin.Context) {
		// repository := postgres.NewDatabaseRepository(dbpool)
		// databaseUsecase := usecase.NewUserUsecase(repository)
		// c.JSON(http.StatusOK, test{
		// 	test: "test",
		// })

		log.Println("I am sampleFunc")
		c.JSON(http.StatusOK, gin.H{"message": "HELLO!"})
	})
}