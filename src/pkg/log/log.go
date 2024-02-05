package log

import (
	"context"

	"go.uber.org/zap"
)

type ctxLogger struct{}

func ContextWithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, logger)
}

func LoggerFromContext(ctx context.Context) *zap.Logger {
	logger := ctx.Value(ctxLogger{}).(*zap.Logger)
	return logger
}
