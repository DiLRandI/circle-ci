package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/DiLRandI/circle-ci/cmd"
	"github.com/DiLRandI/circle-ci/internal/helper"
	"github.com/lmittmann/tint"
)

func main() {
	logger := slog.New(tint.NewHandler(os.Stdout, nil))
	logger = logger.With("app", "circle-ci")

	ctx := context.Background()
	ctx = context.WithValue(ctx, helper.ContextLoggerKey, logger)

	cmd.Execute(ctx)
}
