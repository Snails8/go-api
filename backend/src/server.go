package main

import (
	"fmt"
	"net/http"
)

// http://localhost:7001/
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":7007", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}