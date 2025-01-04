package handler

import (
	"fmt"
	"net/http"
	middleware "api/middleware"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello, World!")
}