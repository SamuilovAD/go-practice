package main

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	ID             int
	ProcessingTime time.Duration
}

func main() {
	orders := []Order{
		{ID: 1, ProcessingTime: 1 * time.Second},
		{ID: 2, ProcessingTime: 3 * time.Second},
		{ID: 3, ProcessingTime: 500 * time.Millisecond},
	}
	var wg sync.WaitGroup
	for _, order := range orders {
		wg.Add(1)
		go handleOrder(order, &wg)
	}
	wg.Wait()
}

func handleOrder(order Order, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case <-time.After(order.ProcessingTime):
		fmt.Println("Order", order.ID, "handled")
	case <-time.After(2 * time.Second):
		fmt.Println("Order", order.ID, "too late")
	}
}
