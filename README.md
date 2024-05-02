# RequestID

[![Run Tests](https://github.com/gin-contrib/requestid/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gin-contrib/requestid/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/requestid/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/requestid)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-contrib/requestid)](https://goreportcard.com/report/github.com/gin-contrib/requestid)
[![GoDoc](https://godoc.org/github.com/gin-contrib/requestid?status.svg)](https://godoc.org/github.com/gin-contrib/requestid)
[![Join the chat at https://gitter.im/gin-gonic/gin](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

Request ID middleware for Gin Framework. Adds an indentifier to the response using the `X-Request-ID` header. Passes the `X-Request-ID` value back to the caller if it's sent in the request headers.

## Usage

### Start using it

Download and install it.

```sh
go get github.com/gin-contrib/requestid
```

Import it in your code, then use it:

```go
import "github.com/gin-contrib/requestid"
```

## Config

define your custom generator function:

```go
func main() {

  r := gin.New()

  r.Use(
    requestid.New(
      requestid.WithGenerator(func() string {
        return "test"
      }),
      requestid.WithCustomHeaderStrKey("your-customer-key"),
    ),
  )

  // Example ping request.
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
  })

  // Listen and Server in 0.0.0.0:8080
  r.Run(":8080")
}
```

## Example

```go
package main

import (
  "fmt"
  "net/http"
  "time"

  "github.com/gin-contrib/requestid"
  "github.com/gin-gonic/gin"
)

func main() {

  r := gin.New()

  r.Use(requestid.New())

  // Example ping request.
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
  })

  // Listen and Server in 0.0.0.0:8080
  r.Run(":8080")
}
```

How to get the request identifier:

```go
// Example / request.
r.GET("/", func(c *gin.Context) {
  c.String(http.StatusOK, "id:"+requestid.Get(c))
})
```

You can also get the request identifier from response header:

```sh
> curl -i "http://127.0.0.1:8080"

HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
X-Request-ID: 77966910-3912-4193-9b74-267491c51700
Content-Length: 39

id:77966910-3912-4193-9b74-267491c51700
```

When http request with custom identifier, gin server return the custom identifier in response header.

```sh
> curl -i -H "X-Request-ID:test" "http://127.0.0.1:8080"

HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
X-Request-Id: 1688221042
Content-Length: 13

id:test
```
