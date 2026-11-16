package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	started := false
	var wg sync.WaitGroup
	wg.Add(3)
	for id := 1; id <= 3; id++ {
		go func(id int) {
			defer wg.Done()
			mu.Lock()
			for !started {
				cond.Wait()
			}
			mu.Unlock()
			fmt.Printf("worker %d started\n", id)
		}(id)
	}
	mu.Lock()
	started = true
	cond.Broadcast()
	mu.Unlock()
	wg.Wait()
}
