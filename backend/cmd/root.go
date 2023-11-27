package cmd

import (
	"github.com/spf13/cobra"
)

func NewPlugin() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "Pimview",
		Short: "Pimview - Media Service",
	}

	rootCmd.AddCommand(RunPlugin())

	return rootCmd
}
