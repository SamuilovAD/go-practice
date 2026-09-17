package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("237"))
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	dict := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var result []string
	var backtrack func(numberIndex int, currentStr string)
	backtrack = func(numberIndex int, currentStr string) {
		if numberIndex == len([]rune(digits)) {
			result = append(result, currentStr)
			return
		}
		letters := dict[string(digits[numberIndex])]
		for _, letter := range letters {
			backtrack(numberIndex+1, currentStr+string(letter))
		}
	}
	backtrack(0, "")

	return result
}
