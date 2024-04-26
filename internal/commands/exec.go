package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var inputUrl string
var numRequests int
var concurrency int

var benchmarkCmd = &cobra.Command{
	Use:   "exec",
	Short: "Run HTTP benchmark",
	Long:  `Run benchmark against a specified URL.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("⚙️ Running benchmark for URL: %s \n✅ Number of Requests: %d \n⌛️ Concurrency: %d\n", inputUrl, numRequests, concurrency)
		runBenchmark(inputUrl, numRequests, concurrency)
	},
}

func init() {
	rootCmd.AddCommand(benchmarkCmd)
	benchmarkCmd.Flags().StringVarP(&inputUrl, "url", "u", "", "URL to benchmark")
	benchmarkCmd.Flags().IntVarP(&numRequests, "requests", "r", 100, "Number of requests to make")
	benchmarkCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 10, "Concurrency level (number of concurrent requests)")
}
