package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n). Extra space: O(1).
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}

func main() {
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2}
	c := &ListNode{Val: 3}
	d := &ListNode{Val: 4}

	a.Next = b
	b.Next = c
	c.Next = d

	// List: 1 -> 2 -> 3 -> 4 -> nil.
	fmt.Println("Without cycle:", hasCycle(a)) // false

	// Connect the last node back to node 2.
	d.Next = b

	fmt.Println("With cycle:", hasCycle(a)) // true
}