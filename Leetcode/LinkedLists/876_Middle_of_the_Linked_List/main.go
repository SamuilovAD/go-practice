package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNodeBruteForce(head *ListNode) *ListNode {
	nodes := []*ListNode{}
	current := head

	for current != nil {
		nodes = append(nodes, current)
		current = current.Next
	}

	return nodes[len(nodes)/2]
}
func middleNodeFastSlowPointers(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	middle := middleNodeBruteForce(head)

	for middle != nil {
		fmt.Printf("%d \n", middle.Val)
		middle = middle.Next
	}

	middle2 := middleNodeFastSlowPointers(head)

	for middle2 != nil {
		fmt.Printf("%d \n", middle2.Val)
		middle2 = middle2.Next
	}
}
