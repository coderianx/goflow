// Package main demonstrates how to reach the underlying *http.Request
// through goflow's Ctx, for example to read query parameters or headers.
//
// Run:
//
//	go run ./examples/request
//
// Then request http://localhost:8080/search?q=goflow.
package main

import (
	"net/http"

	"github.com/coderianx/goflow"
)

func main() {
	// Register a handler that only matches GET requests to "/search".
	http.HandleFunc("GET /search", func(w http.ResponseWriter, r *http.Request) {
		// Wrap the request in a goflow context.
		ctx := goflow.Context(w, r)

		// Ctx.Request exposes the full request, so every helper from the
		// standard library keeps working.
		query := ctx.Request.URL.Query().Get("q")

		// Echo the query back as plain text with status 200 OK.
		ctx.SendString(http.StatusOK, "search query: "+query)
	})

	// Start the server on port 8080. The error is ignored for brevity.
	http.ListenAndServe(":8080", nil)
}
