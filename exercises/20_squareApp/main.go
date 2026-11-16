package main

import (
	"fmt"
	"sync"
)

func main() {
	a, b := 3, 4
	results := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		results <- a * a
	}()
	go func() {
		defer wg.Done()
		results <- b * b
	}()
	go func() {
		wg.Wait()
		close(results)
	}()
	sum := 0
	for result := range results {
		sum += result
	}
	fmt.Printf("Sum: %v", sum)
}
