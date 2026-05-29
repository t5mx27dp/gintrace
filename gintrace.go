package gintrace

import (
	"github.com/gin-gonic/gin"
	"github.com/t5mx27dp/trace"
)

type config struct {
	traceIDKey string
	headerKey  string
}

func New(opts ...Option) gin.HandlerFunc {
	cfg := &config{}

	cfg.traceIDKey = "TraceID"
	cfg.headerKey = "X-Trace-ID"

	for _, opt := range opts {
		opt(cfg)
	}

	return func(c *gin.Context) {
		ctx := c.Request.Context()

		val := c.GetHeader(cfg.headerKey)
		if val == "" {
			ctx = trace.Trace(ctx)
		} else {
			ctx = trace.TraceWithValue(ctx, val)
		}

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
