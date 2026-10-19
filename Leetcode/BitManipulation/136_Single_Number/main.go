package main

import "fmt"

func main() {
    fmt.Println(singleNumber([]int{4,1,2,1,2}))
}

func singleNumber(nums []int) int {
    resultNumber := 0
    for _, num := range nums{
        resultNumber ^= num
    }

    return resultNumber
}