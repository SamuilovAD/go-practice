package main

import "fmt"

func main() {
	input := []int{5, 2, 8, 1, 9, 3, 12}
	bubbleSort(input)
	fmt.Println(input)
}

/*
*
Task:
Implement bubble sort
*/
func bubbleSort(inputArray []int) {
	n := len(inputArray)
	for i := 0; i < n-1; i++ {
		isSwapped := false
		for j := 0; j < n-1-i; j++ {
			if inputArray[j] > inputArray[j+1] {
				inputArray[j], inputArray[j+1] = inputArray[j+1], inputArray[j]
				isSwapped = true
			}
		}
		if !isSwapped {
			break
		}
	}
}
