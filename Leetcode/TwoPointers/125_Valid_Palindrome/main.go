package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println("|", "A man, a plan, a canal: Panama", "|", isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println("|", "race a car", "|", isPalindrome("race a car"))
	fmt.Println("|", " ", "|", isPalindrome(" "))
}
func isPalindrome(s string) bool {
	filteredRunes := make([]rune, 0, len(s))
	for _, ru := range s {
		if unicode.IsLetter(ru) || unicode.IsNumber(ru) {
			filteredRunes = append(filteredRunes, unicode.ToLower(ru))
		}
	}
	leftIndex := 0
	rightIndex := len(filteredRunes) - 1
	for leftIndex < rightIndex {
		if filteredRunes[leftIndex] != filteredRunes[rightIndex] {
			return false
		}
		leftIndex++
		rightIndex--
	}

	return true
}
