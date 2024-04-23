package commands

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mybenchmarktool",
	Short: "A powerful HTTP benchmarking tool",
	Long:  `A CLI tool to benchmark HTTP endpoints.`,
}

func Execute() error {
	return rootCmd.Execute()
}
