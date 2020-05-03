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

	// Example ping request.
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, requestid.Get(c))
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
