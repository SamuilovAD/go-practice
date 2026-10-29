package main

import "fmt"

func main() {
	fmt.Printf("Is isomorfic. %t", isIsomorphic("egg","add"))
}

func isIsomorphic(s string, t string) bool {
    sRunes := []rune(s)
    tRunes := []rune(t)

    if len(sRunes) != len(tRunes) {
        return false
    }

    mappingST := make(map[rune]rune)
    mappingTS := make(map[rune]rune)

    for i, c1 := range sRunes {
        c2 := tRunes[i]

        mappedST, existsST := mappingST[c1]
        mappedTS, existsTS := mappingTS[c2]

        if !existsST && !existsTS {
            mappingST[c1] = c2
            mappingTS[c2] = c1
        } else if !existsST || !existsTS ||
            mappedST != c2 || mappedTS != c1 {
            return false
        }
    }

    return true
}