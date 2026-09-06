package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Create dummy node to handle edge cases (like removing head)
	dummy := &ListNode{Next: head}

	slow := dummy
	fast := dummy

	// Move fast pointer n+1 steps ahead
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	// Move both pointers until fast reaches the end
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// Remove the target node
	slow.Next = slow.Next.Next

	return dummy.Next
}

func main() {
	listNode5 := ListNode{
		Val:  5,
		Next: nil,
	}
	listNode4 := ListNode{
		Val:  4,
		Next: &listNode5,
	}
	listNode3 := ListNode{
		Val:  3,
		Next: &listNode4,
	}
	listNode2 := ListNode{
		Val:  2,
		Next: &listNode3,
	}
	head := ListNode{
		Val:  1,
		Next: &listNode2,
	}
	removeNthFromEnd(&head, 4)
}
