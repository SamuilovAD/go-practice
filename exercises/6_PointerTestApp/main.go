package main

import "fmt"

/*
*
Task:
Create a program that demonstrates how Go slices share the same underlying array.
*/
func main() {
	sliceTest := []int{1, 2, 3, 4, 5}
	sliceTest2 := sliceTest[2:4]
	fmt.Println("sliceTest 1", sliceTest)
	fmt.Println("sliceTest 2", sliceTest2)
	for key, value := range sliceTest {
		sliceTest[key] = value * 2
	}
	fmt.Println("sliceTest 1", sliceTest)
	fmt.Println("sliceTest 2", sliceTest2)
}
