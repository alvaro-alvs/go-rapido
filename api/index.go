package handler

import (
	"fmt"
	"net/http"
	"go-rapido/api/middleware"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	middleware.LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}