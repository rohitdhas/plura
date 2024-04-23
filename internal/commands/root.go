package commands

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "plura",
	Short: "A powerful HTTP benchmarking tool",
	Long: `
	_____  _                
	|  __ \| |
	| |__) | |_   _ _ __ __ _
	|  ___/| | | | | '__/ _  |
	| |    | | |_| | | | (_| |
	|_|    |_|\__,_|_|  \__,_|
							   
	A CLI tool to benchmark HTTP endpoints. 🚀`,
}

func Execute() error {
	return rootCmd.Execute()
}
