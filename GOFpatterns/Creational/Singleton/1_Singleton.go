package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println(getInstance().data)
}

type singleton struct {
	data string
}

var instance *singleton
var once sync.Once

func getInstance() *singleton {
	once.Do(func() {
		instance = &singleton{data: "Data initialized"}
	})
	return instance
}
