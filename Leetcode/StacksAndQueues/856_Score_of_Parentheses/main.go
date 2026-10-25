package main

import "fmt"

func main() {
	fmt.Printf("Parentheses score: %d \n", scoreOfParentheses("(())"))
	fmt.Printf("Parentheses score: %d \n", scoreOfParentheses("(()())"))
	fmt.Printf("Parentheses score: %d \n", scoreOfParentheses("((())())"))
}

func scoreOfParentheses(s string) int {
	stack := Stack{}
	score := 0
	for _, ch := range s {
		if ch == '(' {
			stack.Push(score)
			score = 0
		} else {
			score = stack.Pop() + max(score*2, 1)
		}
	}
	return score
}

type Stack []int

func (stack *Stack) Push(number int) {
	*stack = append(*stack, number)
}

func (stack *Stack) Pop() int {
	number := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]

	return number
}
