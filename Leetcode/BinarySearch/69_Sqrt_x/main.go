package main

import "fmt"

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left := 2
	right := x / 2

	for left <= right {
		pivot := left + (right-left)/2
		num := pivot * pivot

		if num == x {
			return pivot
		}

		if num > x {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}

	return right
}

func main() {
	for _, x := range []int{0, 1, 2, 4, 8, 9, 15, 16, 25, 26} {
		fmt.Printf("mySqrt(%d) = %d\n", x, mySqrt(x))
	}
}