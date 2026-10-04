# goflow

[![Go Reference](https://pkg.go.dev/badge/github.com/coderianx/goflow.svg)](https://pkg.go.dev/github.com/coderianx/goflow)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A small web framework built on top of the Go standard library's
[`net/http`](https://pkg.go.dev/net/http).

## Features

- A request context ([`Ctx`](https://pkg.go.dev/github.com/coderianx/goflow#Ctx))
  that wraps `http.ResponseWriter` and `*http.Request`.
- Helpers for sending plain-text (`SendString`), JSON (`SendJSON`) and HTML
  template (`SendHTML`) responses.
- Redirects with any HTTP status code (`Redirect`).
- Path parameter access (`Param`, `ParamInt`) for routes registered with
  [`http.ServeMux`](https://pkg.go.dev/net/http#ServeMux) patterns.
- Query parameter access (`Query`, `QueryInt`).
- A logger middleware (`middleware.Logger`) that logs the HTTP method, path
  and duration of every request.
- Zero third-party dependencies.

## Installation

```sh
go get github.com/coderianx/goflow
```

## Usage

```go
package main

import (
	"net/http"

	"github.com/coderianx/goflow"
)

func main() {
	http.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		ctx := goflow.Context(w, r)

		ctx.SendString(http.StatusOK, "user: "+ctx.Param("id"))
	})

	http.ListenAndServe(":8080", nil)
}
```

## Examples

Runnable examples live in the [`examples`](examples) directory:

| Example | Description | Run |
| --- | --- | --- |
| [`main.go`](examples/main.go) | Plain-text "Hello World" response | `go run ./examples` |
| [`html`](examples/html) | Rendering an HTML template with `SendHTML` | `go run ./examples/html` |
| [`json`](examples/json) | JSON response with `SendJSON` | `go run ./examples/json` |
| [`middleware`](examples/middleware) | Request logging with `middleware.Logger` | `go run ./examples/middleware` |
| [`pathparams`](examples/pathparams) | Path parameters with `Param` and `ParamInt` | `go run ./examples/pathparams` |
| [`redirect`](examples/redirect) | Redirecting requests with `Redirect` | `go run ./examples/redirect` |
| [`request`](examples/request) | Query parameters with `Query`/`QueryInt` and the underlying `*http.Request` | `go run ./examples/request` |

## API

### `func Context(w http.ResponseWriter, r *http.Request) *Ctx`

Creates a new request context bound to the given response writer and request.

### `func (c *Ctx) Param(name string) string`

Returns the value of the path parameter captured by the route pattern.

### `func (c *Ctx) ParamInt(name string) (int, error)`

Returns the value of the path parameter parsed as an `int`. It returns an
error when the parameter is missing or is not a valid integer.

### `func (c *Ctx) Query(name string) string`

Returns the first value of the named query parameter, or an empty string when
the parameter is absent.

### `func (c *Ctx) QueryInt(name string) (int, error)`

Returns the first value of the named query parameter parsed as an `int`. It
returns an error when the parameter is missing or is not a valid integer.

### `func (c *Ctx) SendString(status int, data string)`

Writes a plain-text response with the given status code.

### `func (c *Ctx) SendJSON(status int, data any)`

Encodes `data` as JSON and writes it with the given status code.

### `func (c *Ctx) SendHTML(path string)`

Parses the HTML template file at `path`, executes it with a `nil` data value
and writes the result with the `text/html; charset=utf-8` content type. The
path is resolved relative to the process working directory. `SendHTML` panics
when the template cannot be parsed.

### `func (c *Ctx) Redirect(status int, url string)`

Sends an HTTP redirect to `url` with the given status code, for example
`http.StatusFound` (302) or `http.StatusMovedPermanently` (301).

### `func middleware.Logger(next http.Handler) http.Handler`

Wraps an `http.Handler` and logs the method, path and duration of every
request once the handler completes.

## License

Released under the [MIT License](LICENSE).
