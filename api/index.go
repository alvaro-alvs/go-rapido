package handler

import (
	"go-rapido/middleware"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	path := r.URL.Path

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		http.
	}

	middleware.LoggingMiddleware(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				http.ServeFile(w, r, path)
			},
		),
	)
}
