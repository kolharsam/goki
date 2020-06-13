package serv

import (
	"net/http"
	"time"

	"github.com/kolharsam/goki/server/logger"
)

// LoggingMiddleware is used to log requests and other info
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method

		// TODO: Logging must be improved...it would be nice to have
		// only the most necessary information being logged
		logger.LogRequest(method, r, time.Now().String())
		next.ServeHTTP(w, r)
		logger.LogInfo(time.Now().String(), w)
	})
}

// EnforceJSONPayloadMiddleware enforces that the request that we get is
// of the Content-Type of `application/json`
func EnforceJSONPayloadMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			next.ServeHTTP(w, r)
			return
		}

		if r.Method == http.MethodPost {
			contentHeader := r.Header.Get("Content-Type")

			if contentHeader != "application/json" {
				http.Error(w, "Malformed Content-Type header", http.StatusBadRequest)
				return
			}
			next.ServeHTTP(w, r)
		}
	})
}
