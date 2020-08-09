package requestid

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const headerXRequestID = "X-Request-ID"

// Config defines the config for RequestID middleware
type Config struct {
	// Generator defines a function to generate an ID.
	// Optional. Default: func() string {
	//   return uuid.New().String()
	// }
	Generator func() string
}

// New initializes the RequestID middleware.
func New(config ...Config) gin.HandlerFunc {
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	}

	// Set config default values
	if cfg.Generator == nil {
		cfg.Generator = func() string {
			return uuid.New().String()
		}
	}

	return func(c *gin.Context) {
		// Get id from request
		rid := c.GetHeader(headerXRequestID)
		if rid == "" {
			rid = cfg.Generator()
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
