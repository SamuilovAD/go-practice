package main

import (
	"fmt"
	"sync"
)

/*
*
Task:
Implement a program in which two goroutines output numbers in turn — odd and even.
*/
func main() {
	maxNumber := 10
	oddChan := make(chan int)
	evenChan := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go printOdd(oddChan, evenChan, &wg)
	go printEven(oddChan, evenChan, &wg, maxNumber)
	oddChan <- 1
	wg.Wait()
}

func printOdd(
	oddChan chan int,
	evenChan chan int,
	waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for {
		number, ok := <-oddChan
		if !ok {
			return
		}
		fmt.Printf("Odd: %d\n", number)
		evenChan <- number + 1
	}
}
func printEven(
	oddChan chan int,
	evenChan chan int,
	waitGroup *sync.WaitGroup,
	max int) {
	defer waitGroup.Done()
	for {
		number, ok := <-evenChan
		if !ok {
			return
		}
		if number > max {
			close(oddChan)
			return
		}
		fmt.Printf("Even: %d\n", number)
		oddChan <- number + 1
	}
}
