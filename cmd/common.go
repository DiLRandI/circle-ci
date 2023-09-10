package cmd

import (
	"context"
	"log"
	"log/slog"

	"github.com/DiLRandI/circle-ci/internal/cmderror"
)

const (
	contextLoggerKey = "ctx_logger"
	circleCiAPIKey   = "api_key"

	envCircleCiAPIKey = "CIRCLE_CI_API_KEY"
)

func tokenFromContext(ctx context.Context) (string, error) {
	token, ok := ctx.Value(circleCiAPIKey).(*string)
	log.Println(token)
	log.Println(ok)
	if !ok || *token == "" {
		return "", cmderror.CircleCiAPIKeyNotFoundError
	}

	return *token, nil
}

func loggerFromContext(ctx context.Context) (*slog.Logger, error) {
	logger, ok := ctx.Value(contextLoggerKey).(*slog.Logger)
	if !ok {
		return nil, cmderror.LoggerNotFoundError
	}

	return logger, nil
}
