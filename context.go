// Package goflow provides a small web framework built on top of the Go
// standard library's net/http package.
//
// Each incoming request is wrapped in a [Ctx], which exposes the underlying
// [http.ResponseWriter] and [http.Request] together with helpers for writing
// plain-text and JSON responses.
package goflow

import "net/http"

// Ctx represents the context of a single HTTP request/response cycle.
// A new Ctx is created for every request that is handled.
type Ctx struct {
	// Writer is the response writer used to send the HTTP response.
	Writer http.ResponseWriter

	// Request is the incoming HTTP request being handled.
	Request *http.Request
}

// Context returns a new [Ctx] bound to the given [http.ResponseWriter] and
// [http.Request]. It is typically called at the beginning of an HTTP handler.
func Context(
	w http.ResponseWriter,
	r *http.Request,
) *Ctx {
	return &Ctx{
		Writer:  w,
		Request: r,
	}
}
