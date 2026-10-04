// Package goflow provides a small web framework built on top of the Go
// standard library's net/http package.
//
// Each incoming request is wrapped in a [Ctx], which exposes the underlying
// [http.ResponseWriter] and [http.Request] together with helpers for writing
// plain-text and JSON responses.
package goflow

import (
	"net/http"
	"strconv"
)

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

// Param returns the value of the path parameter with the given name.
//
// It delegates to [http.Request.PathValue], so it works with the wildcard
// patterns registered on an [http.ServeMux], such as "GET /users/{id}".
func (c *Ctx) Param(name string) string {
	// The captured values are stored on the request by http.ServeMux before
	// the handler is called.
	return c.Request.PathValue(name)
}

func (c *Ctx) ParamInt(name string) (int, error) {
	result, err := strconv.Atoi(
		c.Request.PathValue(name),
	)

	return result, err

}

func (c *Ctx) Query(name string) string {
	return c.Request.URL.Query().Get(name)
}

func (c *Ctx) QueryInt(name string) (int, error) {
	return strconv.Atoi(
		c.Request.URL.Query().Get(name),
	)
}
