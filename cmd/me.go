package cmd

import (
	"github.com/spf13/cobra"
)

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "Get information about the signed in user",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger, err := loggerFromContext(cmd.Context())
		if err != nil {
			return err
		}

		token, err := tokenFromContext(cmd.Context())
		if err != nil {
			return err
		}

		logger.Info("Circle CI API", token)
		return nil
	},
}
