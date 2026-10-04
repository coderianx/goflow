// Package middleware provides reusable net/http middleware that can be
// combined with goflow handlers.
package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logger returns middleware that logs the HTTP method, path and duration of
// every request after the wrapped handler has finished.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Record the start time so the duration can be measured later.
		start := time.Now()

		// Hand the request over to the next handler in the chain.
		next.ServeHTTP(w, r)

		// Calculate how long the handler took.
		duration := time.Since(start)

		// Print the request line with ANSI colors: method in cyan, path in
		// yellow and duration in green.
		log.Printf(
			"\033[36m%s\033[0m \033[33m%s\033[0m \033[32m%s\033[0m",
			r.Method,
			r.URL.Path,
			duration,
		)
	})
}
