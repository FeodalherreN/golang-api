package middleware

import (
	"net/http"
)

// HeaderMiddleware sets default headers on all responses
func HeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(responseWriter, request)
	})
}
