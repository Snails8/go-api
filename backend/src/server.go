package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)





func main() {
	initGinRouter()
}

// use http http://localhost:7001/
func initHttpServer() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":7007", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}

 // use gin.Router listen and server on 0.0.0.0:8080
func initGinRouter() {
	r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello World")
    })
    r.Run()
}

// func main() {
// 	r := gin.Default()
// 	for _, r := range api.NewRouter
// }