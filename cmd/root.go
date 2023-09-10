package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var token string

var rootCmd = &cobra.Command{
	Use:   "circle-ci",
	Short: "A CLI to trigger Circle CI pipelines",
}

func Execute(ctx context.Context) {
	logger, err := loggerFromContext(ctx)
	if err != nil {
		panic(err)
	}

	if err := rootCmd.MarkFlagRequired("circle-ci-api-key"); err != nil {
		logger.Error("required flag is missing", "flag", "circle-ci-api-key", "error", err)
		os.Exit(1)
	}

	ctx = context.WithValue(ctx, circleCiAPIKey, token)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		logger.Error("execution error", "error", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&token, "circle-ci-api-key", "t", "", fmt.Sprintf("Circle CI API key or [%s]", circleCiAPIKey))
	rootCmd.AddCommand(meCmd)
}
