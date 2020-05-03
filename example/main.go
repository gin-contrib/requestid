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
