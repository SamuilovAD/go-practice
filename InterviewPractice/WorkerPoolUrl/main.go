package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL    string
	Status string
	Size   int
	Err    error
}

func worker(
	client *http.Client,
	jobs <-chan string,
	results chan<- Result,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	for url := range jobs {
		result := Result{
			URL: url,
		}
		resp, err := client.Get(url)
		if err != nil {
			result.Err = err
			results <- result
			continue
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			result.Err = err
			results <- result
			continue
		}
		result.Status = resp.Status
		result.Size = len(body)
		results <- result
	}
}

func main() {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://golang.org",
		"https://stackoverflow.com",
		"https://example.com",
	}
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	jobs := make(chan string)
	results := make(chan Result)
	var wg sync.WaitGroup
	const workersCount = 5
	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go worker(
			client,
			jobs,
			results,
			&wg,
		)
	}
	go func() {
		for _, url := range urls {
			jobs <- url
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		if result.Err != nil {
			fmt.Printf(
				"Error: %s, %v \n",
				result.URL,
				result.Err,
			)
		}
		fmt.Printf(
			"Url: %s, Status: %s, Bytes: %d \n",
			result.URL,
			result.Status,
			result.Size,
		)
	}
}
