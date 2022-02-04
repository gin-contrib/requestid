package requestid

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const headerXRequestID = "X-Request-ID"

// Config defines the config for RequestID middleware
type config struct {
	// Generator defines a function to generate an ID.
	// Optional. Default: func() string {
	//   return uuid.New().String()
	// }
	generator Generator
}

// New initializes the RequestID middleware.
func New(opts ...Option) gin.HandlerFunc {
	cfg := &config{
		generator: func() string {
			return uuid.New().String()
		},
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return func(c *gin.Context) {
		// Get id from request
		rid := c.GetHeader(headerXRequestID)
		if rid == "" {
			rid = cfg.generator()
			// Set the id to ensure that the requestid is in the request
			c.Request.Header.Add(headerXRequestID, rid)
		}

		// Set the id to ensure that the requestid is in the response
		c.Header(headerXRequestID, rid)
		c.Next()
	}
}

// Get returns the request identifier
func Get(c *gin.Context) string {
	return c.Writer.Header().Get(headerXRequestID)
}
