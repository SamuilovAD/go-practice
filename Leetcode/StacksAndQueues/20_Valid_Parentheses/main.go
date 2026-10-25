package main

import "fmt"

func main() {
	s := "{{{}}}"
	fmt.Printf(
		"String: %v, is valid paratheses: %v",
		s, isValid(s))
}

var paranthesesMap = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
}

func isValid(s string) bool {
	stack := Stack{}

	for _, ch := range s {
		if _, ok := paranthesesMap[ch]; ok {
			stack.Push(ch)
		} else {
			if stack.IsEmpty() {
				return false
			}

			stackCh := stack.Pop()
			if ch != paranthesesMap[stackCh] {
				return false
			}
		}
	}

	return stack.IsEmpty()
}

type Stack []rune

func (stack *Stack) Push(ch rune) {
	*stack = append(*stack, ch)
}

func (stack *Stack) Pop() rune {
	ch := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]

	return ch
}

func (stack *Stack) IsEmpty() bool {
	return len(*stack) == 0
}
