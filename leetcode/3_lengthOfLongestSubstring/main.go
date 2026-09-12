package main

import "fmt"

// Given a string s, find the length of the longest substring without duplicate Runeacters.
//
// Example 1:
//
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:
//
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:
//
// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
//
// Constraints:
//
// 0 <= s.length <= 5 * 104
// s consists of English letters, digits, symbols and spaces.
func main() {
	fmt.Println(lengthOfLongestStringBruteForce("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

// Time: O(n³) Space Complexity: O(n)
func lengthOfLongestStringBruteForce(s string) int {
	runes := []rune(s)
	longestString := 0
	isUniqueSubstring := func(runes []rune, start int, end int) bool {
		checkedChars := make(map[rune]struct{})
		for i := start; i <= end; i++ {
			if _, ok := checkedChars[runes[i]]; ok {
				return false
			}
			checkedChars[runes[i]] = struct{}{}
		}
		return true
	}
	for i := 0; i < len(runes); i++ {
		for j := i; j < len(runes); j++ {
			if isUniqueSubstring(runes, i, j) {
				if longestString < j-i+1 {
					longestString = j - i + 1
				}
			}
		}
	}
	return longestString
}

// Time: O(n) Space Complexity: O(n)
func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	lastRune := make(map[rune]int)
	leftIndex := 0
	maxLength := 0
	for rightIndex, currentRune := range runes {
		if lastRuneIndex, ok := lastRune[currentRune]; ok && lastRuneIndex >= leftIndex {
			leftIndex = lastRuneIndex + 1
		}
		lastRune[currentRune] = rightIndex
		if currentLength := rightIndex - leftIndex + 1; currentLength > maxLength {
			maxLength = currentLength
		}
	}
	return maxLength
}
