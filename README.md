# RequestID

[![Run Tests](https://github.com/gin-contrib/requestid/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gin-contrib/requestid/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/requestid/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/requestid)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-contrib/requestid)](https://goreportcard.com/report/github.com/gin-contrib/requestid)
[![GoDoc](https://godoc.org/github.com/gin-contrib/requestid?status.svg)](https://godoc.org/github.com/gin-contrib/requestid)
[![Join the chat at https://gitter.im/gin-gonic/gin](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

Request ID middleware for Gin Framework. Adds an indentifier to the response using the `X-Request-ID` header. Passes the `X-Request-ID` value back to the caller if it's sent in the request headers.

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


```mermaid
graph TB
  fe01ce2a7fbac8fafaed7c982a04e229[demo] --> e082d2e93ad95a03eccc88ec3990243a[github.com/json-iterator/go@v1.1.12]
  fe01ce2a7fbac8fafaed7c982a04e229[demo] --> 527c1e693bffacc56634809828702869[github.com/modern-go/concurrent@v0.0.0-20180228061459-e0a39a4cb421]
  fe01ce2a7fbac8fafaed7c982a04e229[demo] --> 631f9901b4e0ea6e08a2a9765c8ed9fb[github.com/modern-go/reflect2@v1.0.2]
  e082d2e93ad95a03eccc88ec3990243a[github.com/json-iterator/go@v1.1.12] --> 2048e7d204077f30102988037755aaa3[github.com/davecgh/go-spew@v1.1.1]
  e082d2e93ad95a03eccc88ec3990243a[github.com/json-iterator/go@v1.1.12] --> 97712d6081a67c995a823384f15d27a6[github.com/google/gofuzz@v1.0.0]
  e082d2e93ad95a03eccc88ec3990243a[github.com/json-iterator/go@v1.1.12] --> 527c1e693bffacc56634809828702869[github.com/modern-go/concurrent@v0.0.0-20180228061459-e0a39a4cb421]
  e082d2e93ad95a03eccc88ec3990243a[github.com/json-iterator/go@v1.1.12] --> 631f9901b4e0ea6e08a2a9765c8ed9fb[github.com/modern-go/reflect2@v1.0.2]
  e082d2e93ad95a03eccc88ec3990243a[github.com/json-iterator/go@v1.1.12] --> fa195f2779d971c124f046ed3e1b0245[github.com/stretchr/testify@v1.3.0]
  fa195f2779d971c124f046ed3e1b0245[github.com/stretchr/testify@v1.3.0] --> b23d389926fc5d61abaaad0c3dc37b1c[github.com/davecgh/go-spew@v1.1.0]
  fa195f2779d971c124f046ed3e1b0245[github.com/stretchr/testify@v1.3.0] --> b256944a298e9d1a6416738704569049[github.com/pmezard/go-difflib@v1.0.0]
  fa195f2779d971c124f046ed3e1b0245[github.com/stretchr/testify@v1.3.0] --> 8e853d38740f4481d09ed4823a906f66[github.com/stretchr/objx@v0.1.0]
```

