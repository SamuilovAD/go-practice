package main

import (
	"fmt"
	"sync"
)

type Task struct {
	name      string
	isHandled bool
}

func worker(id int, tasks <-chan *Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		if !task.isHandled {
			task.name += " is handled"
			task.isHandled = true
			fmt.Printf("Worker %d handled: %s\n", id, task.name)
		}
	}
}

/*
*
Task:
Create a worker pool that processes tasks sent through a channel. Each task has a name and a flag indicating whether it has been handled.
Launch several workers as goroutines; each worker should read tasks from the channel, mark them as handled, and print which worker processed them.
After sending all tasks, close the channel, wait for all workers to finish, and finally print the updated state of every task.
*/
func main() {
	taskList := []*Task{
		{name: "Task1"},
		{name: "Task2"},
		{name: "Task3"},
	}
	taskChan := make(chan *Task)
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, taskChan, &wg)
	}
	for _, t := range taskList {
		taskChan <- t
	}
	close(taskChan)
	wg.Wait()
	fmt.Println("\nFinal task states:")
	for _, t := range taskList {
		fmt.Printf("%s (handled: %v)\n", t.name, t.isHandled)
	}
}
