package commands

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cheggaaa/pb/v3"
)

func runBenchmark(url string, numRequests, concurrency int) {
	client := &http.Client{}

	formattedURL := formatURL(url)

	// Channel to track completion of requests
	complete := make(chan struct{}, numRequests)

	// Start time
	start := time.Now()

	// Create new progress bar
	bar := pb.StartNew(numRequests)
	bar.SetWidth(70)

	// Create workers
	for i := 0; i < concurrency; i++ {
		go func() {
			for j := 0; j < numRequests/concurrency; j++ {
				makeRequest(client, formattedURL, complete, bar)
			}
		}()
	}

	// Wait for all requests to complete
	for i := 0; i < numRequests; i++ {
		<-complete
	}

	// Wait for all requests to complete
	bar.Finish()

	// Calculate elapsed time
	elapsed := time.Since(start)

	// Display benchmark results
	fmt.Printf("\n\t🎉 Benchmark completed in %v\n", elapsed)
	fmt.Printf("\n\t💡 Requests per second: %.2f\n\n", float64(numRequests)/elapsed.Seconds())
}

func makeRequest(client *http.Client, url string, complete chan struct{}, bar *pb.ProgressBar) {
	_, err := client.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
	}

	complete <- struct{}{}

	// Increment progress bar
	bar.Increment()
}

func formatURL(inputURL string) string {
	// Parse the URL
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return ""
	}

	// Remove "www" prefix if present
	parsedURL.Host = removeWWWPrefix(parsedURL.Host)

	// Check if scheme is present, if not, default to "http"
	if parsedURL.Scheme == "" {
		parsedURL.Scheme = "http"
	}

	return parsedURL.String()
}

func removeWWWPrefix(host string) string {
	if len(host) >= 4 && host[:4] == "www." {
		return host[4:]
	}
	return host
}
