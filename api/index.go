package handler

import (
	"fmt"
	"net/http"
	"go-rapido/api/middleware"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	path := r.URL.Path


	if path == "" || path == "/" {
		path = "/index.html"
	}
	
	middleware.LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})).ServeHTTP(w, r)

	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}