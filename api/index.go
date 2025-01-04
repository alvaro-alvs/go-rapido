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

	if (r.Method != "GET") {
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		fmt.println("Olá, Visitante")
	}

	middleware.LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})).ServeHTTP(w, r)
}