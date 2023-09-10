package helper

import (
	"context"
	"log/slog"

	"github.com/DiLRandI/circle-ci/internal/cmderror"
)

type ContextKey string

const (
	ContextLoggerKey = ContextKey("ctx_logger")
	ContextApiKey    = ContextKey("api_key")

	envCircleCiAPIKey = "CIRCLE_CI_API_KEY"
)

func TokenFromContext(ctx context.Context) (string, error) {
	token, ok := ctx.Value(ContextApiKey).(*string)
	if !ok || *token == "" {
		return "", cmderror.CircleCiAPIKeyNotFoundError
	}

	return *token, nil
}

func LoggerFromContext(ctx context.Context) (*slog.Logger, error) {
	logger, ok := ctx.Value(ContextLoggerKey).(*slog.Logger)
	if !ok {
		return nil, cmderror.LoggerNotFoundError
	}

	return logger, nil
}
