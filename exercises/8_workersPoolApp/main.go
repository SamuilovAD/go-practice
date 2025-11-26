package main

import (
	"fmt"
	"sync"
)

type Task struct {
	name      string
	isHandled bool
}

/*
Task:
Run several goroutines that safely mark tasks as handled using a mutex to avoid race conditions, and print the final list of processed tasks.
*/
func main() {
	tasks := []Task{
		{"Task1", false},
		{"Task2", false},
		{"Task3", false},
	}
	var waitGroup sync.WaitGroup
	var mutex sync.Mutex
	for i := 0; i < 3; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			for i := range tasks {
				mutex.Lock()
				if tasks[i].isHandled {
					mutex.Unlock()
					continue
				}
				tasks[i].name = tasks[i].name + " is handled"
				tasks[i].isHandled = true
				mutex.Unlock()
				break
			}
		}()
	}
	waitGroup.Wait()
	fmt.Println(tasks)
}
