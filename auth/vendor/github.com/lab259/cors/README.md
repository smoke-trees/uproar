# Go CORS handler [![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/lab259/cors) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/lab259/cors/master/LICENSE) [![Coverage](http://gocover.io/_badge/github.com/lab259/cors)](http://gocover.io/github.com/lab259/cors)

CORS is a `fasthttp` handler implementing [Cross Origin Resource Sharing W3 specification](http://www.w3.org/TR/cors/) in Golang.

## Getting Started

After installing Go and setting up your [GOPATH](http://golang.org/doc/code.html#GOPATH), create your first `.go` file. We'll call it `server.go`.

```go
package main

import (
    "github.com/lab259/cors"
    "github.com/valyala/fasthttp"
)

func main() {
    // cors.Default() setup the middleware with default options being
    // all origins accepted with simple methods (GET, POST). See
    // documentation below for more options.
    handler := cors.Default().Handler(requestHandler)
    fasthttp.ListenAndServe(":8080", handler)
}

func requestHandler(ctx *fasthttp.RequestCtx) {
    ctx.SetContentType("application/json")
    fmt.Fprintf(ctx, "{\"hello\": \"world\"}")
}
```

Install `cors`:

    go get github.com/lab259/cors

Then run your server:

    go run server.go

The server now runs on `localhost:8080`:

    $ curl -D - -H 'Origin: http://foo.com' http://localhost:8080/
    HTTP/1.1 200 OK
    Access-Control-Allow-Origin: foo.com
    Content-Type: application/json
    Date: Sat, 25 Oct 2014 03:43:57 GMT
    Content-Length: 18

    {"hello": "world"}

### Allow \* With Credentials Security Protection

This library has been modified to avoid a well known security issue when configured with `AllowedOrigins` to `*` and `AllowCredentials` to `true`. Such setup used to make the library reflects the request `Origin` header value, working around a security protection embedded into the standard that makes clients to refuse such configuration. This behavior has been removed with [rs/cors#55](https://github.com/rs/cors/issues/55) and [rs/cors#57](https://github.com/rs/cors/issues/57).

If you depend on this behavior and understand the implications, you can restore it using the `AllowOriginFunc` with `func(origin string) {return true}`.

Please refer to [rs/cors#55](https://github.com/rs/cors/issues/55) for more information about the security implications.

### More Examples

TODO

## Parameters

Parameters are passed to the middleware thru the `cors.New` method as follow:

```go
c := cors.New(cors.Options{
    AllowedOrigins: []string{"http://foo.com", "http://foo.com:8080"},
    AllowCredentials: true,
    // Enable Debugging for testing, consider disabling in production
    Debug: true,
})

// Insert the middleware
handler = c.Handler(handler)
```

- **AllowedOrigins** `[]string`: A list of origins a cross-domain request can be executed from. If the special `*` value is present in the list, all origins will be allowed. An origin may contain a wildcard (`*`) to replace 0 or more characters (i.e.: `http://*.domain.com`). Usage of wildcards implies a small performance penality. Only one wildcard can be used per origin. The default value is `*`.
- **AllowOriginFunc** `func (origin string) bool`: A custom function to validate the origin. It takes the origin as an argument and returns true if allowed, or false otherwise. If this option is set, the content of `AllowedOrigins` is ignored.
- **AllowOriginRequestFunc** `func (r *http.Request origin string) bool`: A custom function to validate the origin. It takes the HTTP Request object and the origin as argument and returns true if allowed or false otherwise. If this option is set, the content of `AllowedOrigins` and `AllowOriginFunc` is ignored
- **AllowedMethods** `[]string`: A list of methods the client is allowed to use with cross-domain requests. Default value is simple methods (`GET` and `POST`).
- **AllowedHeaders** `[]string`: A list of non simple headers the client is allowed to use with cross-domain requests.
- **ExposedHeaders** `[]string`: Indicates which headers are safe to expose to the API of a CORS API specification
- **AllowCredentials** `bool`: Indicates whether the request can include user credentials like cookies, HTTP authentication or client side SSL certificates. The default is `false`.
- **MaxAge** `int`: Indicates how long (in seconds) the results of a preflight request can be cached. The default is `0` which stands for no max age.
- **OptionsPassthrough** `bool`: Instructs preflight to let other potential next handlers to process the `OPTIONS` method. Turn this on if your application handles `OPTIONS`.
- **Debug** `bool`: Debugging flag adds additional output to debug server side CORS issues.

See [API documentation](http://godoc.org/github.com/lab259/cors) for more info.

## Benchmarks

    BenchmarkWithout-8                     200000000   9.45 ns/op  0 B/op     0 allocs/op
    BenchmarkDefault-8                     2000000     646 ns/op   363 B/op   5 allocs/op
    BenchmarkAllowedOrigin-8               2000000     607 ns/op   363 B/op   5 allocs/op
    BenchmarkPreflight-8                   1000000     1322 ns/op  1065 B/op  7 allocs/op
    BenchmarkPreflightHeader-8             1000000     1207 ns/op  1065 B/op  7 allocs/op
    BenchmarkParseHeaderList-8             5000000     338 ns/op   184 B/op   6 allocs/op
    BenchmarkParseHeaderListSingle-8       20000000    85.9 ns/op  32 B/op    3 allocs/op
    BenchmarkParseHeaderListNormalized-8   5000000     312 ns/op   160 B/op   6 allocs/op
    BenchmarkWildcard/match-8              200000000   9.14 ns/op  0 B/op     0 allocs/op
    BenchmarkWildcard/too_short-8          2000000000  1.44 ns/op  0 B/op     0 allocs/op

## Acknowledgments

This is a fork of the incredible [`rs/cors`](https://github.com/rs/cors) package.

## Licenses

All source code is licensed under the [MIT License](https://raw.github.com/lab259/cors/master/LICENSE).
