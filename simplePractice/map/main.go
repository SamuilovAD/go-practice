package main

import (
	"fmt"
	"strconv"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	strings := Map(numbers, func(v int) string {
		return strconv.Itoa(v)
	})
	fmt.Println(strings)
	squares := Map(numbers, func(v int) int {
		return v * v
	})
	fmt.Println(squares)
}

func Map[T any, R any](items []T, fn func(T) R) []R {
	result := make([]R, 0, len(items))
	for _, item := range items {
		result = append(result, fn(item))
	}

	return result
}
