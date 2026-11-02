package requestid

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	headerXRequestID = "X-Request-ID"
	headerKeyContext = "github.com/gin-contrib/requestid.headerKey"
)

// Config defines the config for RequestID middleware
type config struct {
	// Generator defines a function to generate an ID.
	// Optional. Default: func() string {
	//   return uuid.New().String()
	// }
	generator Generator
	headerKey HeaderStrKey
	handler   Handler
}

// New initializes the RequestID middleware.
func New(opts ...Option) gin.HandlerFunc {
	cfg := &config{
		generator: func() string {
			return uuid.New().String()
		},
		headerKey: headerXRequestID,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	headerKey := string(cfg.headerKey)

	return func(c *gin.Context) {
		c.Set(headerKeyContext, headerKey)

		// Get id from request
		rid := c.GetHeader(headerKey)
		if rid == "" {
			rid = cfg.generator()
			c.Request.Header.Add(headerKey, rid)
		}
		if cfg.handler != nil {
			cfg.handler(c, rid)
		}
		// Set the id to ensure that the requestid is in the response
		c.Header(headerKey, rid)
		c.Next()
	}
}

// Get returns the request identifier
func Get(c *gin.Context) string {
	headerKey := c.GetString(headerKeyContext)
	if headerKey == "" {
		headerKey = headerXRequestID
	}
	return c.GetHeader(headerKey)
}
