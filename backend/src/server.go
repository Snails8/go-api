package main

import (
	"context"
	"fmt"
	"go-api/api"
	"go-api/middleware"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	ctx := context.Background()
	
	dbpool, err := pgxpool.Connect(ctx, os.Getenv("DB_DSN"))
	if err != nil {
		fmt.Println("cloud not connect db")
	}

	logger := &middleware.Logger{}

	r := api.NewRouter(dbpool, logger)
	r.Run(":7000")  // docker の published port と衝突しないようにする
}

// ---------  use http http://localhost:7001/ ---------------------------
func initHttpServer() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":7007", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}

// ------------ use gin.Router listen and server on 0.0.0.0:8080 -------------
func initGinRouterSimple() {
	r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello World")
    })
    r.Run()
}

// ----------------- use gin.Router -----------------------------------------
type Route struct {
	name string
	method string
	pattern string
	handlerFunc func(dbpool *pgxpool.Pool, logger *middleware.Logger) gin.HandlerFunc
}
type Routes []Route

var routes = Routes{
	{
		"GetUsers",
		http.MethodGet,
		"/api/users",
		api.GetUsers,
	},
}

func NewRouter(dbpool *pgxpool.Pool, logger *middleware.Logger) *gin.Engine {
	router := gin.Default()
	for _, route := range routes {
		switch route.method {
		case http.MethodGet:
			router.GET(route.pattern, route.handlerFunc(dbpool, logger))
		case http.MethodPost:
			router.POST(route.pattern, route.handlerFunc(dbpool, logger))
		case http.MethodPut:
			router.PUT(route.pattern, route.handlerFunc(dbpool, logger))
		case http.MethodDelete:
			router.DELETE(route.pattern, route.handlerFunc(dbpool, logger))
		}
	}
	
	return router
}

// -------------- gin router 2 -----------------------------------------
// func initGinRouter(dbpool *pgxpool.Pool, logger *middleware.Logger) {
// 	r := gin.Default()

// 	for _, r := range api.NewRouter(dbpool, logger).Route {

// 	}
// }