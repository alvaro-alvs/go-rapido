package main

import (
	
	"go-rapido/middleware"
)

func main() {
	middleware.LoggingMiddleware(http.HandlerFunc(handler.Handler)).ServeHTTP()
}