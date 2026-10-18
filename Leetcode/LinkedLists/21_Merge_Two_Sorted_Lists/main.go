package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var head *ListNode
	if list1.Val < list2.Val {
		head = list1
		list1 = list1.Next
	} else {
		head = list2
		list2 = list2.Next
	}

	head.Next = mergeTwoLists(list1, list2)
	return head
}

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{
		Val: 3, Next: &ListNode{Val: 5},
	}}
	list2 := &ListNode{Val: 2, Next: &ListNode{
		Val: 4, Next: &ListNode{Val: 6},
	}}

	merged := mergeTwoLists(list1, list2)

	for node := merged; node != nil; node = node.Next {
		fmt.Print(node.Val)
		if node.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}