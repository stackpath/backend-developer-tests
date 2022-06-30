package middlewares

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// request ID constants
const (
	XRequestIDHeader = "X-Request-ID"
	XRequestIDKey    = "RequestID"
)

// RequestID middleware sets the tracing request ID to context
// If the request ID is present in the header, it will use it or generate one
// if it is not present
func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqID := ctx.Request.Header.Get(XRequestIDHeader)
		if reqID == "" {
			reqID = uuid.NewV4().String()
		}
		ctx.Request.Header.Set(XRequestIDHeader, reqID)
		ctx.Writer.Header().Set(XRequestIDHeader, reqID)
		ctx.Set(XRequestIDKey, reqID)
		ctx.Next()
	}
}
