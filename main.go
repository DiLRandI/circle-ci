package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/DiLRandI/circle-ci/cmd"
)

func main() {
	var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	ctx := context.Background()
	ctx = context.WithValue(ctx, "ctx_logger", logger)
	cmd.Execute(ctx)
}
