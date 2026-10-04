// Package main demonstrates how to render an HTML template file with
// goflow's SendHTML helper.
//
// Run this example from the repository root:
//
//	go run ./examples/html
//
// Then open http://localhost:8080/hello in a browser.
package main

import (
	"net/http"

	"github.com/coderianx/goflow"
)

// templatePath points at the template file that lives next to this program.
// SendHTML resolves the path relative to the working directory, so the
// example is run from the repository root.
const templatePath = "examples/html/index.html"

func main() {
	// Register a handler for "/hello".
	http.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		// Wrap the request in a goflow context.
		ctx := goflow.Context(w, r)

		// Render the template and send it as HTML with status 200 OK.
		ctx.SendHTML(templatePath)
	})

	// Start the server on port 8080. The error is ignored for brevity.
	http.ListenAndServe(":8080", nil)
}
