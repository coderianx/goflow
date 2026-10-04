// Package main demonstrates how to use the logger middleware with goflow.
//
// Run:
//
//	go run ./examples/middleware
//
// Then request http://localhost:8080/hello and watch the server output.
package main

import (
	"net/http"

	"github.com/coderianx/goflow"
	"github.com/coderianx/goflow/middleware"
)

func main() {
	// Create a mux that holds the routes.
	mux := http.NewServeMux()

	// Register a simple route.
	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		// Wrap the request in a goflow context.
		ctx := goflow.Context(w, r)

		ctx.SendString(http.StatusOK, "Hello from a logged request")
	})

	// Serve the mux through the logger middleware so every request is logged
	// once it finishes. The error is ignored for brevity.
	http.ListenAndServe(":8080", middleware.Logger(mux))
}
