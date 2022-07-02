package middlewares

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	ContextRedactHeaders = "RedactHeaders"
	Redacted             = "[---redacted---]"
)

// Variables for redacting
var (
	redactHeadersFields = map[string]bool{
		"Authorization": true,
	}
)

// Logging is an application logging gin middleware
func Logging() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		// initialize application logger
		_, logger := log.Get(ctx)
		defer logger.Sync()

		// add request ID to logger
		if reqID := ctx.GetString(XRequestIDKey); reqID != "" {
			logger = logger.With(zap.String(XRequestIDKey, reqID))
		}

		// log request start time
		logger.With(zap.Time("start_time", start)).Info("Request Start")

		// set redact fields to context
		ctx.Set(ContextRedactHeaders, redactHeadersFields)

		// add more fields to log here
		logFields := []zap.Field{}
		logger = logger.With(logFields...)

		// log request information if logging is set to debug level
		if logger.Core().Enabled(zapcore.DebugLevel) {
			headers, err := redactHeaders(ctx)
			if err != nil {
				logger.With(zap.String("error", err.Error())).Warn("Failed reading request headers")
			}
			logger.With(zap.String("headers", headers)).Debug("Request headers")
		}

		// add logger to context
		logger.With(zap.Time("timestamp", time.Now()))
		ctx.Set(log.ContextLoggerKey, logger)
		ctx.Next()

		// log request end time
		end := start.Add(time.Since(start))
		logger.With(zap.Time("end_time", end)).Info("Request End")
	}
}

// redactHeaders extract and return header in JSON format after
// redacting headers defined in context by ContextRedactHeaders field
func redactHeaders(c *gin.Context) (string, error) {
	redactHeaders := c.Value(ContextRedactHeaders)
	rhs, ok := redactHeaders.(map[string]bool)
	if !ok {
		rhs = map[string]bool{}
	}

	headers := map[string]string{}
	for h := range c.Request.Header {
		if rhs[h] {
			headers[h] = Redacted
		} else {
			headers[h] = c.Request.Header.Get(h)
		}
	}
	jsonHeaders, err := json.Marshal(headers)
	if err != nil {
		return "", err
	}
	return string(jsonHeaders), nil
}
