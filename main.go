/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/DiLRandI/circle-ci/cmd"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	ctx := context.Background()
	ctx = context.WithValue(ctx, "ctx_logger", logger)

	cmd.Execute(ctx)
}
