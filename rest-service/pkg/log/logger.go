package log

import (
	"context"
	"os"

	"go.uber.org/zap"
)

// Logging constants
const (
	EnvLogLevel        = "LOG_LEVEL"
	EnvDevelopment     = "LOG_DEVELOPMENT"
	SamplingInitial    = 100
	SamplingThereafter = 100

	ContextLoggerKey = "ContextLogger"
)

// Logger instance for logging
type Logger struct {
	*zap.Logger
}

func init() {
	// Initialize defaults
	if os.Getenv(EnvLogLevel) == "" {
		if err := os.Setenv(EnvLogLevel, zap.DebugLevel.String()); err != nil {
			panic(err.Error())
		}
	}
}

// New initializes a new logger
func New(opts ...zap.Option) (*Logger, error) {
	level, err := zap.ParseAtomicLevel(os.Getenv(EnvLogLevel))
	if err != nil {
		return nil, err
	}

	// Start with production encoder and update as required
	ec := zap.NewProductionEncoderConfig()
	zcfg := zap.Config{
		Level:       level,
		Development: os.Getenv(EnvDevelopment) == "true",
		Sampling: &zap.SamplingConfig{
			Initial:    SamplingInitial,
			Thereafter: SamplingThereafter,
		},
		Encoding:         "json",
		EncoderConfig:    ec,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, err := zcfg.Build(opts...)
	if err != nil {
		return nil, err
	}
	defer logger.Sync()
	return &Logger{logger}, nil
}

// Get returns logger if already initialized or
// creates and returns logger if not initialized
func Get(ctx context.Context, opts ...zap.Option) (context.Context, *Logger) {
	if logger, ok := ctx.Value(ContextLoggerKey).(*Logger); ok {
		return ctx, logger
	}
	// Create and return new logger
	logger, err := New(opts...)
	if err != nil {
		panic(err.Error())
	}

	ctx = context.WithValue(ctx, ContextLoggerKey, logger)
	return ctx, logger
}

// With returns a new logger with additional fields
func (l *Logger) With(fields ...zap.Field) *Logger {
	return &Logger{l.Logger.With(fields...)}
}
