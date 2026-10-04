package goflow

import (
	"encoding/json"
	"io"
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
