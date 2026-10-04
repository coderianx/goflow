// Package main demonstrates how to read path parameters with goflow.
//
// goflow uses the pattern matching of the standard library's http.ServeMux:
// register a pattern such as "GET /users/{id}/posts/{postID}" and read the
// captured values with Ctx.Param. Ctx.ParamInt additionally parses a value as
// an integer.
//
// Run:
//
//	go run ./examples/pathparams
//
// Then request http://localhost:8080/users/7/posts/99.
package main

import (
	"fmt"
	"net/http"

	"github.com/coderianx/goflow"
)

func main() {
	// Curly braces mark the dynamic segments of the route; http.ServeMux
	// captures them by name.
	http.HandleFunc("GET /users/{id}/posts/{postID}", func(w http.ResponseWriter, r *http.Request) {
		// Wrap the request in a goflow context.
		ctx := goflow.Context(w, r)

		// Param returns the captured value as a string, exactly as it
		// appeared in the request path.
		postID := ctx.Param("postID")

		// ParamInt parses the captured value as an integer.
		userID, err := ctx.ParamInt("id")
		if err != nil {
			// The route pattern accepts any text, so reject non-numeric
			// input instead of failing later.
			ctx.SendString(http.StatusBadRequest, "invalid user id")
			return
		}

		// Send the parsed values back as plain text with status 200 OK.
		ctx.SendString(
			http.StatusOK,
			fmt.Sprintf("user=%d, post=%s", userID, postID),
		)
	})

	// Start the server on port 8080. The error is ignored for brevity.
	http.ListenAndServe(":8080", nil)
}
