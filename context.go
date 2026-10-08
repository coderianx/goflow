// Package goflow provides a small web framework built on top of the Go
// standard library's net/http package.
//
// Each incoming request is wrapped in a [Ctx], which exposes the underlying
// [http.ResponseWriter] and [http.Request] together with helpers for writing
// plain-text, JSON and HTML responses, sending redirects and reading path and
// query parameters.
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

// ParamInt returns the value of the path parameter with the given name
// parsed as an integer. It returns an error if the parameter is missing or
// cannot be parsed with [strconv.Atoi].
func (c *Ctx) ParamInt(name string) (int, error) {
	result, err := strconv.Atoi(
		c.Request.PathValue(name),
	)

	return result, err
}

// Query returns the first value of the query parameter with the given name,
// or an empty string when the parameter is absent.
func (c *Ctx) Query(name string) string {
	return c.Request.URL.Query().Get(name)
}

// QueryDefault returns the first value of the query parameter with the given
// name, or the fallback value when the parameter is absent or empty.
func (c *Ctx) QueryDefault(name, fallback string) string {
	value := c.Query(name)

	if value == "" {
		return fallback
	}

	return value
}

// QueryInt returns the first value of the query parameter with the given name
// parsed as an integer. It returns an error if the parameter is missing or
// cannot be parsed with [strconv.Atoi].
func (c *Ctx) QueryInt(name string) (int, error) {
	return strconv.Atoi(
		c.Request.URL.Query().Get(name),
	)
}

// QueryIntDefault returns the first value of the query parameter with the
// given name parsed as an integer, or the fallback value when the parameter
// is absent or empty. It returns an error if the parameter is present but
// cannot be parsed with [strconv.Atoi].
func (c *Ctx) QueryIntDefault(name string, fallback int) (int, error) {
	value := c.Query(name)

	if value == "" {
		return fallback, nil
	}

	val, err := strconv.Atoi(value)

	return val, err
}

// QueryInt64 returns the first value of the query parameter with the given name
// parsed as an int64. It returns an error if the parameter is missing or
// cannot be parsed with [strconv.ParseInt].
func (c *Ctx) QueryInt64(name string) (int64, error) {
	val, err := strconv.ParseInt(
		c.Request.URL.Query().Get(name), 10, 64,
	)

	return val, err
}

// QueryInt64Default returns the first value of the query parameter with the
// given name parsed as an int64, or the fallback value when the parameter
// is absent or empty. It returns an error if the parameter is present but
// cannot be parsed with [strconv.ParseInt].
func (c *Ctx) QueryInt64Default(name string, fallback int64) (int64, error) {
	value := c.Query(name)

	if value == "" {
		return fallback, nil
	}

	val, err := strconv.ParseInt(
		value, 10, 64,
	)

	return val, err
}

// QueryFloat returns the first value of the query parameter with the given name
// parsed as a float64. It returns an error if the parameter is missing or
// cannot be parsed with [strconv.ParseFloat].
func (c *Ctx) QueryFloat(name string) (float64, error) {
	val, err := strconv.ParseFloat(
		c.Request.URL.Query().Get(name), 64,
	)

	return val, err
}

// QueryFloatDefault returns the first value of the query parameter with the
// given name parsed as a float64, or the fallback value when the parameter
// is absent or empty. It returns an error if the parameter is present but
// cannot be parsed with [strconv.ParseFloat].
func (c *Ctx) QueryFloatDefault(name string, fallback float64) (float64, error) {
	value := c.Query(name)

	if value == "" {
		return fallback, nil
	}

	val, err := strconv.ParseFloat(
		value, 64,
	)

	return val, err
}

// QueryBool returns the first value of the query parameter with the given name
// parsed as a boolean. It returns an error if the parameter is missing or
// cannot be parsed with [strconv.ParseBool].
func (c *Ctx) QueryBool(name string) (bool, error) {
	val, err := strconv.ParseBool(
		c.Request.URL.Query().Get(name),
	)

	return val, err
}

// QueryBoolDefault returns the first value of the query parameter with the
// given name parsed as a boolean, or the fallback value when the parameter
// is absent or empty. It returns an error if the parameter is present but
// cannot be parsed with [strconv.ParseBool].
func (c *Ctx) QueryBoolDefault(name string, fallback bool) (bool, error) {
	value := c.Query(name)

	if value == "" {
		return fallback, nil
	}

	val, err := strconv.ParseBool(value)

	return val, err
}

// Header returns the value of the given header name. It returns an empty
// string when the header is not present. The name is case-insensitive.
func (c *Ctx) Header(name string) string {
	return c.Request.Header.Get(name)
}

// HeaderSet sets the given header name to the specified value. It replaces
// any existing values for that header. The name is case-insensitive.
func (c *Ctx) HeaderSet(key, value string) {
	c.Writer.Header().Set(key, value)
}
