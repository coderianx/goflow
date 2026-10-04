package goflow

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
)

// SendString writes a plain-text response with the given HTTP status code.
// It sets the Content-Type header to "text/plain; charset=utf-8".
func (c *Ctx) SendString(status int, data string) {
	// Announce the response as UTF-8 plain text.
	c.Writer.Header().Set(
		"Content-Type", "text/plain; charset=utf-8",
	)

	// Write the HTTP status code, for example http.StatusOK.
	c.Writer.WriteHeader(status)

	// Write the response body. The error is ignored because the status code
	// and headers have already been sent.
	io.WriteString(c.Writer, data)
}

// SendJSON writes data to the response as JSON with the given HTTP status
// code. It sets the Content-Type header to "application/json".
func (c *Ctx) SendJSON(status int, data any) {
	// Announce the response as JSON.
	c.Writer.Header().Set(
		"Content-Type", "application/json",
	)

	// Write the HTTP status code, for example http.StatusCreated.
	c.Writer.WriteHeader(status)

	// Encode data into the response body. Encoding errors are ignored
	// because the status code has already been sent and cannot be changed.
	_ = json.NewEncoder(c.Writer).Encode(data)
}

// SendHTML parses the HTML template file at path, executes it with a nil data
// value and writes the rendered result to the response. It sets the
// Content-Type header to "text/html; charset=utf-8".
//
// The path is resolved relative to the process working directory. Templates
// are parsed with [html/template], so rendered output is HTML-escaped.
// SendHTML panics when the template cannot be parsed, because it relies on
// [template.Must].
func (c *Ctx) SendHTML(path string) {
	// Announce the response as HTML.
	c.Writer.Header().Set(
		"Content-Type", "text/html; charset=utf-8",
	)

	// Parse the template file. Must panics on malformed templates, which
	// surfaces programming errors early during development.
	tmpl := template.Must(template.ParseFiles(path))

	// Render the template. The error is ignored because the headers have
	// already been sent by the time execution can fail.
	tmpl.Execute(c.Writer, nil)
}

// Redirect sends an HTTP redirect to url with the given status code, such as
// [http.StatusFound] or [http.StatusMovedPermanently].
//
// It delegates to [http.Redirect].
func (c *Ctx) Redirect(status int, url string) {
	// http.Redirect sets the Location header and writes the status code
	// together with a short HTML body for clients that do not follow
	// redirects automatically.
	http.Redirect(c.Writer, c.Request, url, status)
}
