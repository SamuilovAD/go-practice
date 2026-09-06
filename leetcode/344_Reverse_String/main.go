package main

import "fmt"

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(string(s))
	reverseString(s)
	fmt.Println(string(s))
}

func reverseString(s []byte) {
	leftIndex := 0
	rightIndex := len(s) - 1
	for leftIndex < rightIndex {
		s[leftIndex], s[rightIndex] = s[rightIndex], s[leftIndex]
		leftIndex++
		rightIndex--
	}
}
