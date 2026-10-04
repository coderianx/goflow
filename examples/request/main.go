// Package main demonstrates how to read query parameters and how to reach the
// underlying *http.Request through goflow's Ctx.
//
// Run:
//
//	go run ./examples/request
//
// Then request http://localhost:8080/search?q=goflow&page=2.
package main

import (
	"fmt"
	"net/http"

	"github.com/coderianx/goflow"
)

func main() {
	// Register a handler that only matches GET requests to "/search".
	http.HandleFunc("GET /search", func(w http.ResponseWriter, r *http.Request) {
		// Wrap the request in a goflow context.
		ctx := goflow.Context(w, r)

		// Query returns a single query parameter as a string.
		query := ctx.Query("q")

		// QueryInt does the same for numeric parameters. "page" is
		// optional, so fall back to 1 when it is missing or invalid.
		page, err := ctx.QueryInt("page")
		if err != nil {
			page = 1
		}

		// Ctx.Request still exposes the full request for anything the
		// helpers do not cover, such as reading headers.
		userAgent := ctx.Request.UserAgent()

		// Echo the values back as plain text with status 200 OK.
		ctx.SendString(
			http.StatusOK,
			fmt.Sprintf(
				"search query: %s, page: %d, user agent: %s",
				query, page, userAgent,
			),
		)
	})

	// Start the server on port 8080. The error is ignored for brevity.
	http.ListenAndServe(":8080", nil)
}
