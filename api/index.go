package handler

import (
	"fmt"
	"net/http"
	"api/middleware"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	middleware.LoggingMiddleware(http.HandlerFunc(handler)).ServeHTTP(w, r)

	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}