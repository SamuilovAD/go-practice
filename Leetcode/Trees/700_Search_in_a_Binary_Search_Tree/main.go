package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(h), where h is the height of the tree.
// Auxiliary space: O(h) for the recursive call stack.
// Both are O(log n) for a balanced tree and O(n) in the worst case.
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if val == root.Val {
		return root
	}
	if val < root.Val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 7},
	}

	for i, val := range []int{2, 5} {
		node := searchBST(root, val)
		fmt.Printf("Iteration #%d \n", i)
		if node == nil {
			fmt.Printf("%d not found\n", val)
			continue
		}
		fmt.Printf("Found node with value %d\n", node.Val)
	}
}
