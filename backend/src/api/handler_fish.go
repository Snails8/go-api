package api

import (
	"go-api/internal/fish/repository/postgres"
	"go-api/internal/fish/usecase"
	"go-api/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func GetFishes(dbpool *pgxpool.Pool, logger *middleware.Logger) gin.HandlerFunc {
	return gin.HandlerFunc(func (c *gin.Context) {
		r := postgres.NewFishRepository(dbpool)
		u := usecase.NewFishUsecase(r)

		fishesList := u.GetFishes(c, logger)
		fishes := make([]Fish, 0)
		for _, fish := range fishesList {
			fish := Fish{
				Id: fish.Id,
				Name: fish.Name,
			}
			fishes = append(fishes, fish)
		}

		c.JSON(http.StatusOK, GetFishesResponse{
			Fishes: fishes,
		})
	})
}