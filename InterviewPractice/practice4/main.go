package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	text := `
Go is simple, and Go is fast.
Python is simple too!
GO is great; Python is popular.
`
	limit := 3
	fmt.Printf("%v", TopWords(text, limit))
}

type WordCount struct {
	Word  string
	Count int
}

func TopWords(text string, limit int) []WordCount {
	if text == "" || limit <= 0 {
		return []WordCount{}

	}
	splitedWordsMap := strings.Fields(text)
	counterMap := make(map[string]int)
	for _, word := range splitedWordsMap {
		word = strings.ToLower(word)
		word = strings.Trim(word, ".,!?:;")
		if word == "" {
			continue
		}
		counterMap[word] += 1
	}
	result := make([]WordCount, 0, len(counterMap))
	for filtredWord, wordCount := range counterMap {
		result = append(result, WordCount{
			Word:  filtredWord,
			Count: wordCount,
		})
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Count != result[j].Count {
			return result[i].Count > result[j].Count
		}
		return result[i].Word < result[j].Word
	})
	if len(result) > limit {
		result = result[0:limit]
	}
	return result
}
