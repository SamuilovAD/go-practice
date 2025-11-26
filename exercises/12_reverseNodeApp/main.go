package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	reversed := reverseList(node)

	for reversed != nil {
		fmt.Print(reversed.Val, " ")
		reversed = reversed.Next
	}
}
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head

	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}

	return prev
}
