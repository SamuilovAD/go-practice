package main

import "fmt"

func main() {
	fmt.Println(isAnagram("anagram", "naagarm"))
	fmt.Println(isAnagram("anagram1", "n2agarm"))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	results := make(map[rune]int, len([]rune(s)))
	for _, ch := range s {
		results[ch]++
	}
	for _, ch := range t {
		results[ch]--
	}
	for _, count := range results {
		if count != 0 {
			return false
		}
	}
	return true
}
