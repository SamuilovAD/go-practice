// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := []Job{
		{ID: 10, Value: 5},
		{ID: 20, Value: 3},
		{ID: 30, Value: 8},
		{ID: 40, Value: 2},
	}
	workers := 2
	processedJobs := ProcessJobs(jobs, workers)
	fmt.Print(processedJobs)
}

type Job struct {
	ID    int
	Value int
}

type Result struct {
	JobID int
	Value int
}
type JobIndexed struct {
	job   Job
	index int
}
type ResultIndexed struct {
	result Result
	index  int
}

func ProcessJobs(jobs []Job, workers int) []Result {
	if workers <= 0 || len(jobs) <= 0 {
		return []Result{}
	}
	var wg sync.WaitGroup
	requestJobsIndexedCh := make(chan JobIndexed, workers)
	responseJobsIndexedCh := make(chan ResultIndexed, workers)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go workerHandler(&wg, requestJobsIndexedCh, responseJobsIndexedCh)
	}
	go func() {
		for jobIndex, job := range jobs {
			requestJobsIndexedCh <- JobIndexed{
				job:   job,
				index: jobIndex,
			}
		}
		close(requestJobsIndexedCh)
	}()
	go func() {
		wg.Wait()
		close(responseJobsIndexedCh)
	}()
	results := make([]Result, len(jobs))
	for responseJobIndexed := range responseJobsIndexedCh {
		results[responseJobIndexed.index] = responseJobIndexed.result
	}

	return results
}
func workerHandler(wg *sync.WaitGroup, requestJobsIndexedCh chan JobIndexed, resultJobsIndexedCh chan ResultIndexed) {
	defer wg.Done()
	for requestJobIndexed := range requestJobsIndexedCh {
		result := ResultIndexed{
			index: requestJobIndexed.index,
			result: Result{
				JobID: requestJobIndexed.job.ID,
				Value: requestJobIndexed.job.Value * requestJobIndexed.job.Value,
			},
		}
		resultJobsIndexedCh <- result
	}
}
