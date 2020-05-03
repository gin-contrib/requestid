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
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
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
	"regexp"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

var (
	rxURL = regexp.MustCompile(`^/regexp\d*`)
)

func main() {

	r := gin.New()

	r.Use(requestid.New())

	// Example ping request.
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
```
