package cmd

import (
	"fmt"

	"github.com/DiLRandI/circle-ci/internal/helper"
	"github.com/spf13/cobra"
)

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "Get information about the signed in user",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger, err := helper.LoggerFromContext(cmd.Context())
		if err != nil {
			return fmt.Errorf("failed to get logger from context: %w", err)
		}

		token, err := helper.TokenFromContext(cmd.Context())
		if err != nil {
			return fmt.Errorf("failed to get token from context: %w", err)
		}

		logger.Info("Circle CI API", "token", token)

		return nil
	},
}
