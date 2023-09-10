/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/DiLRandI/circle-ci/cmd"
	"github.com/DiLRandI/circle-ci/internal/helper"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	ctx := context.Background()
	ctx = context.WithValue(ctx, helper.ContextLoggerKey, logger)

	cmd.Execute(ctx)
}
