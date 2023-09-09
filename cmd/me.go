package cmd

import (
	"log"

	"github.com/DiLRandI/circle-ci/internal/cmderror"

	"github.com/spf13/cobra"
)

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "Get information about the signed in user",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger, ok := cmd.Context().Value(contextLoggerKey).(*log.Logger)
		if !ok {
			return cmderror.LoggerNotFoundError
		}

		logger.Println("me called")

		return nil
	},
}
