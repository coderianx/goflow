// Package main demonstrates how to send a JSON response with goflow's
// SendJSON helper.
//
// Run:
//
//	go run ./examples/json
//
// Then request http://localhost:8080/user.
package main

import (
	"net/http"

	"github.com/coderianx/goflow"
)

// user is the payload returned by the example endpoint. The struct tags
// control the field names used in the JSON output.
type user struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Register a handler for "/user". Method patterns such as "GET /user"
	// are supported by http.ServeMux since Go 1.22.
	http.HandleFunc("GET /user", func(w http.ResponseWriter, r *http.Request) {
		// Wrap the request in a goflow context.
		ctx := goflow.Context(w, r)

		// Build the payload and send it as JSON with status 200 OK.
		ctx.SendJSON(http.StatusOK, user{
			ID:    1,
			Name:  "Ada Lovelace",
			Email: "ada@example.com",
		})
	})

	// Start the server on port 8080. The error is ignored for brevity.
	http.ListenAndServe(":8080", nil)
}
