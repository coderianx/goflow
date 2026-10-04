// Package main demonstrates how to read path parameters with goflow.
//
// goflow uses the pattern matching of the standard library's http.ServeMux:
// register a pattern such as "GET /users/{id}/posts/{postID}" and read the
// captured values with Ctx.Param.
//
// Run:
//
//	go run ./examples/pathparams
//
// Then request http://localhost:8080/users/7/posts/99.
package main

import (
	"net/http"

	"github.com/coderianx/goflow"
)

func main() {
	// Curly braces mark the dynamic segments of the route; http.ServeMux
	// captures them by name.
	http.HandleFunc("GET /users/{id}/posts/{postID}", func(w http.ResponseWriter, r *http.Request) {
		// Wrap the request in a goflow context.
		ctx := goflow.Context(w, r)

		// Read the values captured from the request path.
		userID := ctx.Param("id")
		postID := ctx.Param("postID")

		// Send them back as plain text with status 200 OK.
		ctx.SendString(http.StatusOK, "user="+userID+", post="+postID)
	})

	// Start the server on port 8080. The error is ignored for brevity.
	http.ListenAndServe(":8080", nil)
}
