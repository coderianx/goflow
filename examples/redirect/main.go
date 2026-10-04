// Package main demonstrates how to send HTTP redirects with goflow's
// Redirect helper.
//
// Run:
//
//	go run ./examples/redirect
//
// Then request http://localhost:8080/old and watch the client follow the
// redirect to /new.
package main

import (
	"net/http"

	"github.com/coderianx/goflow"
)

func main() {
	// Permanently redirect "/old" to "/new". Use http.StatusFound instead
	// when the redirect is only temporary.
	http.HandleFunc("GET /old", func(w http.ResponseWriter, r *http.Request) {
		// Wrap the request in a goflow context.
		ctx := goflow.Context(w, r)

		// 301 tells clients and search engines to use the new location.
		ctx.Redirect(http.StatusMovedPermanently, "/new")
	})

	// The redirect target.
	http.HandleFunc("GET /new", func(w http.ResponseWriter, r *http.Request) {
		// Wrap the request in a goflow context.
		ctx := goflow.Context(w, r)

		// Send a plain-text response with status 200 OK.
		ctx.SendString(http.StatusOK, "You have been redirected")
	})

	// Start the server on port 8080. The error is ignored for brevity.
	http.ListenAndServe(":8080", nil)
}
