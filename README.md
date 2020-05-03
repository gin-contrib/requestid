# RequestID

Request ID middleware for Gin Framework. Adds an indentifier to the response using the `X-Request-ID` header

## Config

define your custom generator function:

```go
func main() {

	r := gin.New()

	r.Use(requestid.New(requestid.Config{
		Generator: func() string {
			return "test"
		},
	}))

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
	// Example ping request.
	r.GET("/ping", func(c *gin.Context) {
    id := requestid.Get(c)
		c.String(http.StatusOK, id)
	})
```
