package main

import (
	"fmt"
	"sync"
)

func multiplyByTwo(input <-chan int, output chan<- int, wg *sync.WaitGroup) {
	defer close(output)
	defer wg.Done()

	for num := range input {
		output <- num * 2
	}
}
func filterLessThan10(input <-chan int, output chan<- int, wg *sync.WaitGroup) {
	defer close(output)
	defer wg.Done()

	for num := range input {
		if num >= 10 {
			output <- num
		}
	}
}
func printResults(input <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range input {
		fmt.Println("Result:", num)
	}
}

/*
*
Task:
Create a pipeline of three goroutines that process numbers step by step:
1. The first goroutine reads numbers from a channel and multiplies each one by 2.
2. The second goroutine receives these results and forwards only the values that are ≥ 10.
3. The third goroutine reads the remaining numbers and prints them.
Send the numbers `1, 4, 5, 6, 10, 15` into the pipeline.
Make sure the channels are closed properly and all goroutines finish their work.
*/
func main() {
	var wg sync.WaitGroup
	stage1 := make(chan int)
	stage2 := make(chan int)
	stage3 := make(chan int)
	wg.Add(3)
	go multiplyByTwo(stage1, stage2, &wg)
	go filterLessThan10(stage2, stage3, &wg)
	go printResults(stage3, &wg)
	go func() {
		defer close(stage1)
		inputNumbers := []int{1, 4, 5, 6, 10, 15}
		for _, num := range inputNumbers {
			stage1 <- num
		}
	}()

	wg.Wait()
}
