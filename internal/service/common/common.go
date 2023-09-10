package common

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/DiLRandI/circle-ci/internal/helper"
)

const (
	Endpoint = "https://circleci.com/api/v2/"
)

func GetLoggerAndToken(ctx context.Context) (*slog.Logger, string, error) {
	logger, err := helper.LoggerFromContext(ctx)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get logger from context: %w", err)
	}

	token, err := helper.TokenFromContext(ctx)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get token from context: %w", err)
	}

	return logger, token, nil
}
