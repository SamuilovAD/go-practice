package main

import "fmt"

func main() {
	fmt.Println(findAnagrams("cbaebabacd","abc"))
}

func findAnagrams(s string, p string) []int {
    if len(s) < len(p) {
        return []int{}
    }
    var source, target [26]int
    for i := 0; i < len(p); i++ {
        source[s[i] - 'a']++ 
        target[p[i] - 'a']++
    }
    result := []int{}
    if source == target {
        result = append(result, 0)
    }
    for i := len(p); i < len(s); i++ {
        source[s[i] - 'a']++
        source[s[i - len(p)] - 'a']--
        if target == source {
            result = append(result, i - len(p) + 1)
        }
    }
    return result
}
