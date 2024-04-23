package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var url string
var method string
var body string

var benchmarkCmd = &cobra.Command{
	Use:   "exec",
	Short: "Run HTTP benchmark",
	Long:  `Run benchmark against a specified URL.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Execute benchmark with provided parameters
		fmt.Printf("Running benchmark for URL: %s, Method: %s, Body: %s\n", url, method, body)
		// Call benchmarking function with provided parameters
	},
}

func init() {
	rootCmd.AddCommand(benchmarkCmd)
	benchmarkCmd.Flags().StringVarP(&url, "url", "u", "", "URL to benchmark")
	benchmarkCmd.Flags().StringVarP(&method, "method", "m", "GET", "HTTP method (GET, POST, PUT, etc.)")
	benchmarkCmd.Flags().StringVarP(&body, "body", "b", "", "Request body")
}
