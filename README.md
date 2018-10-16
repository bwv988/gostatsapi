# gostatsapi - A simple REST API example in Go

This is my first attempt at learning the Go programming language.

For this purpose, I've decided to implement a very simple REST server which exposes an API to calculate summary statistics.

## Overview

- Uses the awesome `chi` framework for routing & request handling in Go.
- Why `chi`? Because it's fast and fully compatible with `net/http`.
- The calculation results are obtained through the `gonum` package.

### Implementation notes, REST & best practices

There are some good guides on how to go about this. Here are the principles I'm attempting to follow:

- API is versioned. Allows for a graceful deprecation of clients.
- Use of proper HTTP verbs.
- For isolation and maintainability, each route has its own package.
- Basic integration testing via CURL.

### Installing and running

Run the following from the project's top-level directory:

```bash
go get -u github.com/go-chi/chi
go get -u github.com/go-chi/render
go go get gonum.org/v1/gonum/...

cd src

go run main.go

# In a separate terminal:

cd test
./test_median.sh
```

## TODOs

- [] Checking of boundary conditions and better error handling.
- [] Implement dependency management for 3pp libs.
- [] Improve documentation.
- [] Use proper golang test framework instead of CURL.
- [] Refactor package structure.

## References
- https://itnext.io/structuring-a-production-grade-rest-api-in-golang-c0229b3feedc
- https://github.com/go-chi/chi
- https://attilaolah.eu/2014/09/10/json-and-struct-composition-in-go/
- http://www.gonum.org/
