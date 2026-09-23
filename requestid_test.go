package requestid

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const testXRequestID = "test-request-id"

func emptySuccessResponse(c *gin.Context) {
	c.String(http.StatusOK, "")
}

func Test_RequestID_CreateNew(t *testing.T) {
	r := gin.New()
	r.Use(New())
	r.GET("/", emptySuccessResponse)

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Header().Get(headerXRequestID))
}

func Test_RequestID_PassThru(t *testing.T) {
	r := gin.New()
	r.Use(New())
	r.GET("/", emptySuccessResponse)

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	req.Header.Set(headerXRequestID, testXRequestID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testXRequestID, w.Header().Get(headerXRequestID))
}

func TestRequestIDWithCustomID(t *testing.T) {
	r := gin.New()
	r.Use(
		New(
			WithGenerator(func() string {
				return testXRequestID
			}),
		),
	)
	r.GET("/", emptySuccessResponse)

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testXRequestID, w.Header().Get(headerXRequestID))
}

func TestRequestIDWithCustomHeaderKey(t *testing.T) {
	r := gin.New()
	r.Use(
		New(
			WithCustomHeaderStrKey("customKey"),
		),
	)
	r.GET("/", emptySuccessResponse)

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	req.Header.Set("customKey", testXRequestID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testXRequestID, w.Header().Get("customKey"))
}

func TestRequestIDWithCustomHeaderKeysAreIndependent(t *testing.T) {
	const (
		firstHeader  = "X-Request-ID-First"
		secondHeader = "X-Request-ID-Second"
		firstID      = "first-request-id"
		secondID     = "second-request-id"
	)

	firstRouter := gin.New()
	firstRouter.Use(New(WithCustomHeaderStrKey(firstHeader)))
	firstRouter.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, Get(c))
	})

	secondRouter := gin.New()
	secondRouter.Use(New(WithCustomHeaderStrKey(secondHeader)))
	secondRouter.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, Get(c))
	})

	tests := []struct {
		name      string
		router    *gin.Engine
		headerKey string
		requestID string
	}{
		{name: "first router", router: firstRouter, headerKey: firstHeader, requestID: firstID},
		{name: "second router", router: secondRouter, headerKey: secondHeader, requestID: secondID},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
			req.Header.Set(tt.headerKey, tt.requestID)

			tt.router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, tt.requestID, w.Header().Get(tt.headerKey))
			assert.Equal(t, tt.requestID, w.Body.String())
		})
	}
}

func TestNewConcurrent(t *testing.T) {
	const goroutines = 16

	start := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := range goroutines {
		headerKey := HeaderStrKey(fmt.Sprintf("X-Request-ID-%d", i))
		go func() {
			defer wg.Done()
			<-start

			for range 100 {
				New(WithCustomHeaderStrKey(headerKey))
			}
		}()
	}

	close(start)
	wg.Wait()
}

func TestRequestIDWithHandler(t *testing.T) {
	r := gin.New()
	called := false
	r.Use(
		New(
			WithHandler(func(c *gin.Context, requestID string) {
				called = true
				assert.Equal(t, testXRequestID, requestID)
			}),
		),
	)

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	req.Header.Set("X-Request-ID", testXRequestID)
	r.ServeHTTP(w, req)

	assert.True(t, called)
}

func TestRequestIDIsAttachedToRequestHeaders(t *testing.T) {
	r := gin.New()

	r.Use(New())

	r.GET("/", func(c *gin.Context) {
		result := c.GetHeader("X-Request-ID")
		assert.NotEmpty(t, result)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	r.ServeHTTP(w, req)
}

func TestRequestIDNotNilAfterGinCopy(t *testing.T) {
	r := gin.New()
	r.Use(New())

	r.GET("/", func(c *gin.Context) {
		copy := c.Copy()
		result := Get(copy)
		assert.NotEmpty(t, result)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	r.ServeHTTP(w, req)
}
