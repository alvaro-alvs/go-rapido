package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// LoggingMiddleware é um middleware que registra os detalhes de cada requisição.
func LoggingMiddleware(next http.Handler) http.Handler {
	useColors := os.Getenv("USE_COLORS") == "true" // Verifica se a variável está definida como "true"
	reset := ""
	green := ""
	yellow := ""
	red := ""

	if useColors { // Define as cores apenas se USE_COLORS for "true"
		reset = "\033[0m"
		green = "\033[32m"
		yellow = "\033[33m"
		red = "\033[31m"
	}

	statusColor := green

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Registrar o início da requisição
		start := time.Now()

		// Criar um ResponseWriter personalizado para capturar o status
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		if lrw.statusCode >= 400 && lrw.statusCode < 500 {
			statusColor = yellow
		} else if lrw.statusCode >= 500 {
			statusColor = red
		}

		// Chamar o próximo handler na cadeia
		next.ServeHTTP(lrw, r)

		// Registrar informações da requisição
		fmt.Printf(
			"%s%s%s | Status: %s%d%s | %s | %v\n",
			green, r.Method, reset,
			statusColor, lrw.statusCode, reset,
			r.URL.Path,
			time.Since(start),
		)
	})
}

// loggingResponseWriter é um ResponseWriter que captura o código de status
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captura o código de status e o salva no loggingResponseWriter
func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}