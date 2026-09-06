package main

import (
	"fmt"
	"math"
)

/*
*
Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

Example 1:
Input: x = 123
Output: 321

Example 2:
Input: x = -123
Output: -321
Example 3:

Input: x = 120
Output: 21

Constraints:

-231 <= x <= 231 - 1
*/
func main() {
	fmt.Println(reverseSimple(123))
}

func reverse(x int) int {
	const maxInt = math.MaxInt32
	const minInt = math.MinInt32
	res := 0
	for x != 0 {
		digit := x % 10
		x /= 10
		// overflow check BEFORE multiplication
		if res > maxInt/10 || (res == maxInt/10 && digit > 7) {
			return 0
		}
		if res < minInt/10 || (res == minInt/10 && digit < -8) {
			return 0
		}

		res = res*10 + digit
	}

	return res
}

func reverseSimple(x int) int {
	res := 0
	for x != 0 {
		res = res*10 + x%10
		x /= 10
	}
	return res
}
