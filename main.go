package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/DiLRandI/circle-ci/cmd"
)

func main() {
	var logger = slog.NewLogLogger(slog.NewTextHandler(os.Stdout, nil), slog.LevelInfo)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "ctx_logger", logger)
	cmd.Execute(ctx)
}
