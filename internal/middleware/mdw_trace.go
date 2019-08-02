package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

const (
	TraceID = "centaur_trace_id"
)

//TraceMiddleware 实现了一个跟踪ID中间件
func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetHeader("X-Request-ID")
		if traceID == "" {
			traceID = uuid.NewV4().String()
		}

		c.Set(TraceID, traceID)
		c.Next()
	}
}
