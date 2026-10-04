# goflow

[![Go Reference](https://pkg.go.dev/badge/github.com/coderianx/goflow.svg)](https://pkg.go.dev/github.com/coderianx/goflow)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A small web framework built on top of the Go standard library's
[`net/http`](https://pkg.go.dev/net/http).

## Features

- A request context ([`Ctx`](https://pkg.go.dev/github.com/coderianx/goflow#Ctx))
  that wraps `http.ResponseWriter` and `*http.Request`.
- Helpers for sending plain-text (`SendString`) and JSON (`SendJSON`)
  responses.
- Path parameter access (`Param`) for routes registered with
  [`http.ServeMux`](https://pkg.go.dev/net/http#ServeMux) patterns.
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
| [`json`](examples/json) | JSON response with `SendJSON` | `go run ./examples/json` |
| [`middleware`](examples/middleware) | Request logging with `middleware.Logger` | `go run ./examples/middleware` |
| [`pathparams`](examples/pathparams) | Path parameters with `Param` | `go run ./examples/pathparams` |
| [`request`](examples/request) | Accessing the underlying `*http.Request` | `go run ./examples/request` |

## API

### `func Context(w http.ResponseWriter, r *http.Request) *Ctx`

Creates a new request context bound to the given response writer and request.

### `func (c *Ctx) Param(name string) string`

Returns the value of the path parameter captured by the route pattern.

### `func (c *Ctx) SendString(status int, data string)`

Writes a plain-text response with the given status code.

### `func (c *Ctx) SendJSON(status int, data any)`

Encodes `data` as JSON and writes it with the given status code.

### `func middleware.Logger(next http.Handler) http.Handler`

Wraps an `http.Handler` and logs the method, path and duration of every
request once the handler completes.

## License

Released under the [MIT License](LICENSE).
