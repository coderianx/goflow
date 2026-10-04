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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := goflow.Context(w, r)

		ctx.SendString(200, "Hello World")
	})

	http.ListenAndServe(":8080", nil)
}
```

A runnable version of this example lives in
[`examples/main.go`](examples/main.go):

```sh
go run ./examples
```

## API

### `func Context(w http.ResponseWriter, r *http.Request) *Ctx`

Creates a new request context bound to the given response writer and request.

### `func (c *Ctx) SendString(status int, data string)`

Writes a plain-text response with the given status code.

### `func (c *Ctx) SendJSON(status int, data any)`

Encodes `data` as JSON and writes it with the given status code.

## License

Released under the [MIT License](LICENSE).
