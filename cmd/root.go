/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/DiLRandI/circle-ci/internal/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	token   string
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "circle-ci",
	Short: "A CLI to trigger Circle CI pipelines",
}

func Execute(ctx context.Context) {
	logger, err := helper.LoggerFromContext(ctx)
	if err != nil {
		panic(err)
	}

	if err := rootCmd.MarkFlagRequired("circle-ci-api-key"); err != nil {
		logger.Error("required flag is missing", "flag", "circle-ci-api-key", "error", err)
		os.Exit(1)
	}

	ctx = context.WithValue(ctx, helper.ContextApiKey, token)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		logger.Error("execution error", "error", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.circle-ci.yaml)")
	rootCmd.PersistentFlags().StringVarP(&token, "circle-ci-api-key", "t", "", "Circle CI API key or [CIRCLE_CI_API_KEY]")

	rootCmd.Flags().BoolP("toggle", "", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".circle-ci" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".circle-ci")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
