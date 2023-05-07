package middleware

import (
	"github.com/benbjohnson/clock"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewLoggingMiddleware(logger *zap.Logger, clock clock.Clock) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := clock.Now()

		// Process request
		c.Next()

		// Finish timer and log request
		end := clock.Now()
		latency := end.Sub(start)

		logger.Info("Received request",
			zap.Time("start", start),
			zap.Time("end", end),
			zap.Duration("latency", latency),
			zap.String("path", c.Request.URL.Path),
			zap.Stringer("url", c.Request.URL),
			zap.String("method", c.Request.Method),
			zap.Int("statusCode", c.Writer.Status()),
			zap.Int("responseSize", c.Writer.Size()),
		)
	}
}
