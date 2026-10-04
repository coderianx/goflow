// Package main demonstrates the simplest goflow program: it starts an HTTP
// server on port 8080 and replies "Hello World" to every request.
//
// Run:
//
//	go run ./examples
//
// More focused examples live in the subdirectories of this directory:
// ./json, ./middleware, ./pathparams and ./request.
package main

import (
	"net/http"

	"github.com/coderianx/goflow"
)

func main() {
	// Register a handler for the root path.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Wrap the request in a goflow context.
		ctx := goflow.Context(w, r)

		// Send a plain-text response with status 200 OK.
		ctx.SendString(http.StatusOK, "Hello World")
	})

	// Start the server on port 8080. The error is ignored for brevity.
	http.ListenAndServe(":8080", nil)
}
