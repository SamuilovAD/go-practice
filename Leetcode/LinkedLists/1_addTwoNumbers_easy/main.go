package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return dummy.Next
}

func buildList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, n := range nums {
		current.Next = &ListNode{Val: n}
		current = current.Next
	}
	return dummy.Next
}

func printList(node *ListNode) {
	fmt.Print("[")
	for node != nil {
		fmt.Print(node.Val)
		if node.Next != nil {
			fmt.Print(", ")
		}
		node = node.Next
	}
	fmt.Println("]")
}

func main() {
	l1 := buildList([]int{9, 4, 3}) // 349
	l2 := buildList([]int{5, 6, 8}) // 865

	result := addTwoNumbers(l1, l2)
	printList(result) // [7, 0, 8]
}
