package middleware

import (
	"log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		log.Printf("[%s] %s %s - Remote: %s", r.Method, r.URL.Path, r.Proto, r.RemoteAddr)

		next.ServeHTTP(wrapped, r)

		duration := time.Since(start)
		log.Printf("[%s] %s - Status: %d - Duration: %v", r.Method, r.URL.Path, wrapped.statusCode, duration)
	})
}
