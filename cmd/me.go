package cmd

import (
	"fmt"
	"time"

	httpclient "github.com/DiLRandI/circle-ci/internal/httpClient"
	"github.com/DiLRandI/circle-ci/internal/service/me"
	"github.com/spf13/cobra"
)

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "Get information about the signed in user",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		svc := me.New(httpclient.New(5 * time.Second))
		data, err := svc.GetMe(cmd.Context())
		if err != nil {
			return fmt.Errorf("failed to get me: %w", err)
		}

		data.Print(cmd.OutOrStdout())

		return nil
	},
}
